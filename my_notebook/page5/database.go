import (
  "strconv"
  "strings"
  "sync"
  "time"

  "github.com/ethereum/go-ethereum/log"
  "github.com/ethereum/go-ethereum/metrics"
  "github.com/syndtr/goleveldb/leveldb"
  "github.com/syndtr/goleveldb/leveldb/errors"
  "github.com/syndtr/goleveldb/leveldb/filter"
  "github.com/syndtr/goleveldb/leveldb/iterator"
  "github.com/syndtr/goleveldb/leveldb/opt"
  gometrics "github.com/rcrowley/go-metrics"
)


type LDBDatabase struct {
  fn string 
  db *leveldb.DB  

  getTimer    gometrics.Timer 
  putTimer    gometrics.Timer

  quitLock      sync.Mutex
  quitChan      chan chan error 

  log log.Logger   
}


func NewLDBDatabase(file string, cache int, handles int) (* LDBDatabase, error) {
  logger := log.New("database", file)

  if cache < 16 {
    cache = 16
  } 

  if handles < 16 {
    handles = 16
  }

  logger.info("allocate cache and file handles", "cache", cache, "handles", handles)

  db, err := leveldb.OpenFile(file, &opt.Options{
      OpenFileCacheCapacity:      handles,
      BlockCacheCapacity:         cache / 2 * opt.MiB,
      WriteBuffer:                cache / 4 * opt.MiB,
      Filter:                     filter.NewBloomFilter(10),
    })

  if _, corrupted := err.(*errors.ErrCorrupted); corrupted {
    db, err = leveldb.RecoverFile(file, nil) 
  }

  if err != nil {
    return nil, err 
  }

  return &LDBDatabase {
    fn:   file,
    db:   db,
    log:  logger,
  }, nil

}




func (db *LDBDatabase) Meter(prefix string) {
  // short  circuit metering 
  if !metrics.Enabled {
    return
  }

  db.getTimer = metrics.NewTimer(prefix + "user/gets")

  // create a quit channel for periodic collector and run it
  db.quitLock.Lock()
  db.quitChan = make(chan chan error)       // 初始化这个 quitChan ?????
  db.quitLock.Unlock()

  go db.meter(3 * time.Second)
}








