

type MemDatabase struct {
  db map[string][]byte
  lock sync.RWMutex
}


func NewMemDatabase() (*MemDatabase, error) {
  return &MemDatabase{
    db: make(map[string][]byte),
  }, nil
}


func (db *MemDatabase) Put(key []byte, value []byte) error {
  db.lock.Lock()      // 锁住全部线程,包括读和写

  defer db.lock.Unlock()    // 函数执行完就解锁

  db.db[string(key)] = common.CopyBytes(value)
  return nil
}


func (db *MemDatabase) Has(key []byte) (bool, error) {
  db.lock.Rlock()       // 锁住 读 的线程
  defer db.lock.Unlock()

  _, ok := db.db[string(key)]
  return ok, nil
}



// ======================= batch ======================================

type kv struct {
  k, v []byte 
}

type memBatch struct {
  db *MemDatabase
  writes []kv
  size int
}

func (b *memBatch) Put(key, value []byte) error {
  b.writes = append(b.writes, kv{common.CopyBytes(key), common.CopyBytes(value)})
  b.size += len(value)
  return nil
}


func (b *memBatch) Write() error {
  b.db.lock.Lock()
  defer b.db.lock.Unlock()

  for _, kv := range b.writes {
    b.db.db[string(kv.k)] = kv.v
  }
  return nil

}