package main

import (
  "sync"
  "reflect"
  "fmt"
)

var (
  typeCacheMutex sync.RWMutex                  //读写锁，用来在多线程的时候保护typeCache这个Map
  typeCache= make(map[typekey]*typeinfo) //核心数据结构，保存了类型->编解码器函数
)

type typeinfo struct {
  //存储了编码器和解码器函数
  decoder
  writer
}
type typekey struct {
  reflect.Type
  // the key must include the struct tags because they
  // might generate a different decoder.
  tags
}

func cachedTypeInfo(typ reflect.Type, tags tags) (*typeinfo, error) {
  typeCacheMutex.RLock() //加读锁来保护，
  info := typeCache[typekey{typ, tags}]
  typeCacheMutex.RUnlock()
  if info != nil { //如果成功获取到信息，那么就返回
    return info, nil
  }
  // not in the cache, need to generate info for this type.
  typeCacheMutex.Lock() //否则加写锁 调用cachedTypeInfo1函数创建并返回， 这里需要注意的是在多线程环境下有可能多个线程同时调用到这个地方，所以当你进入cachedTypeInfo1方法的时候需要判断一下是否已经被别的线程先创建成功了。
  defer typeCacheMutex.Unlock()
  return cachedTypeInfo1(typ, tags)
}



func cachedTypeInfo1(typ reflect.Type, tags tags) (*typeinfo, error) {
  key := typekey{typ, tags}
  info := typeCache[key]
  if info != nil {
    // 其他的线程可能已经创建成功了， 那么我们直接获取到信息然后返回
    return info, nil
  }
  // put a dummmy value into the cache before generating.
  // if the generator tries to lookup itself, it will get
  // the dummy value and won't call itself recursively.
  //这个地方首先创建了一个值来填充这个类型的位置，避免遇到一些递归定义的数据类型形成死循环
  typeCache[key] = new(typeinfo)
  info, err := genTypeInfo(typ, tags)
  if err != nil {
    // remove the dummy value if the generator fails
    delete(typeCache, key)
    return nil, err
  }
  *typeCache[key] = *info
  return typeCache[key], err
}

func genTypeInfo(typ reflect.Type, tags tags) (info *typeinfo, err error) {
  info = new(typeinfo)
  if info.decoder, err = makeDecoder(typ, tags); err != nil {
    return nil, err
  }
  if info.writer, err = makeWriter(typ, tags); err != nil {
    return nil, err
  }
  return info, nil
}

// makeWriter creates a writer function for the given type.
func makeWriter(typ reflect.Type, ts tags) (writer, error) {
  kind := typ.Kind()
  switch {
  case typ == rawValueType:
    return writeRawValue, nil
  case typ.Implements(encoderInterface):
    return writeEncoder, nil
  case kind != reflect.Ptr && reflect.PtrTo(typ).Implements(encoderInterface):
    return writeEncoderNoPtr, nil
  case kind == reflect.Interface:
    return writeInterface, nil
  case typ.AssignableTo(reflect.PtrTo(bigInt)):
    return writeBigIntPtr, nil
  case typ.AssignableTo(bigInt):
    return writeBigIntNoPtr, nil
  case isUint(kind):
    return writeUint, nil
  case kind == reflect.Bool:
    return writeBool, nil
  case kind == reflect.String:
    return writeString, nil
  case kind == reflect.Slice && isByte(typ.Elem()):
    return writeBytes, nil
  case kind == reflect.Array && isByte(typ.Elem()):
    return writeByteArray, nil
  case kind == reflect.Slice || kind == reflect.Array:
    return makeSliceWriter(typ, ts)
  case kind == reflect.Struct:
    return makeStructWriter(typ)
  case kind == reflect.Ptr:
    return makePtrWriter(typ)
  default:
    return nil, fmt.Errorf("rlp: type %v is not RLP-serializable", typ)
  }
}





type field struct {
  index int
  info *typeinfo
}

func makeStructWriter(typ reflect.Type) (writer, error) {
  fields, err := structFields(typ)
  if err != nil {
    return nil, err
  }
  writer := func(val reflect.Value, w *encbuf) error {
    lh := w.list()
    for _, f := range fields {
      //f是field结构， f.info是typeinfo的指针， 所以这里其实是调用字段的编码器方法。
      if err := f.info.writer(val.Field(f.index), w); err != nil {
        return err
      }
    }
    w.listEnd(lh)
    return nil
  }
  return writer, nil
}








































