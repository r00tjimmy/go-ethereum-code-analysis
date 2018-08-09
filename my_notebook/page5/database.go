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
}