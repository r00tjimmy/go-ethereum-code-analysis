package ethdb


const idea_batch_size = 100 * 1024


// Putter 接口定义了批量操作和普通操作的写入接口
type Putter interface {
  Put(key []byte, value []byte) error
}



type Database interface {
  Putter      // 属性也可以是另外一个接口
  Get(key []byte) ([]byte, error)
  Has(key []byte) (bool, error)
  Delete(key []byte) error
  Close()
  NewBatch() Batch
}

 

type Batch interface  {
  Putter
  ValueSize() int
  Write() error
}

