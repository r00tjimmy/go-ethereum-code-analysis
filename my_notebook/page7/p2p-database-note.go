
// newNodeDB 创建一个节点数据库用来储存节点信息的
func newNodeDB (path string, version int, self NodeID) (*nodeDB, error) {
  if path == "" {
    return newMemoryNodeDB(self)
  }  

  return newPersistentNodeDB(path, version, self)
}


func newMemoryNodeDB (self NodeID) (*nodeDB, error) {
  db, err := leveldb.Open(storage.NewMemStorage(), nil)
  if err != nil {
    return nil, err
  }  

  return &nodeDB {
    lvl:      db, 
    self:     self,
    quit:     make(chan struct{}),
  }, nil
}


func newPersistentNodeDB(path string, version int, self NodeID) (*nodeDB, error) {
  opts := &opt.Options{ OpenFilesCacheCapacity: 5 }
  db, err := leveldb.OpenFile(path, opts)
  
  if _, iscorrupted := err.(*errors.ErrCorrupoted); iscorrupted {
    db, err = leveldb.RecoverFIle(path, nil)
  }


  if err != nil {
    return nil, err
  }

  currentVer := make([]byte, binary.MaxVarintLen64)
  currentVer = currentVer[:binary.PutVarint(currentVer, int64(version))]
  blob, err := db.Get(nodeDBVersionKey, nil)

  switch err {
  case leveldb.ErrNotFound:
    if err := db.Put(nodeDBVersionKey, currentVer, nil); err != nil {
      db.Close()
      return nil, err
    } 

  case nil:
    if !bytes.Equal(blob, currentVer) {
      db.CLose()
      if err = os.RemoveAll(path); err != nil {
        return nil, err
      } 
      return newPersistentNodeDB(path, version, self)
    }
  }

  return &nodeDB {
    lvl:    db,
    self:   self,
    quit:   make(chan struct{}),
  }, nil

}





func (db *nodeDB) node(id NodeID) *Node {
  blob, err := db.lvl.Get(makeKey(id, nodeDBDiscoverRoot), nil)
  if err != nil {
    return nil
  }

  node := new(Node)
  if err := rlp.DecodeBytes(blob, node); err != nil {
    log.Error("fail to decode node RLP", "err", err)
    return nil
  } 
  node.sha = crypto.Keccak256Hash(node.ID[:])
  return node
}


func (db *nodeDB) updateNode(node *Node) error {
  blob, err := rlp.EncodeToBytes(node)
  if err != nil {
    return err
  }
  return db.lvl.Put(makeKey(node.ID, nodeDBDiscoverRoot), blob. nil)
}



type Node struct {
  IP            net.IP
  UDP, TCP      uint16
  ID            NodeID
  sha           common.Hash 
  contested     bool 
}



func (db *nodeDB) ensureExpirer() {
  db.runner.Do(func() { go db.expirer() })
}



func (db *nodeDB) expirer() {
  tick := time.Tick(nodeDBCleanupCycle)
  for {
    select {
    case <- tick:
      if err := db.expireNodes(); err != nil {
        log.Error("fail to expire nodedb items", "error", err)
      }

    case <- db.quit:
      return
    }
  }
}


func (db *nodeDB) expireNodes() error {
  threshold := time.Now().Add(-nodeDBNodeExpiration)

  it := db.lvl.NewIterator(nil, nil)

  defer it.Release()
  
  for it.Next() {
    id, field := splitKey(it.Key())
    if field != nodeDBDiscoverRoot {
      continue
    }

    if !bytes.Equal(id[:], db.self[:]) {
      if seen := db.lastPong(id); seen.After(threshold) {
        continue
      }
    }

    db.deleteNode(id)
  } 

  return nil 
}




func (db *nodeDB) querySeeds(n int, maxAge time.Duration) []*Node {
  var (
    now     =   time.Now()
    nodes   =   make([]*Node, 0, n)
    it      =   db.lvl.NewIterator(nil, nil)
    id      NodeID 
  )

  defer  it.Release()

  seek:
    for seeks := 0; len(nodes) < n && seeks < n*5; seeks++ {
      ctr := id[0]
      rand.Read(id[:])
      id[0] = ctr + id[0] % 16
      it.Seek(makeKey(id, nodeDBDiscoverRoot))

      n := nextNode(it)
      if n == nil {
        id[0] = 0
        continue seek
      }

      if n.ID == db.self {
        continue seek
      }

      if now.Sub(dn.lastPong(n.ID)) > maxAge {
        continue seek
      }

      for i := range nodes {
        if nodes[i].ID  == n.ID {
          continue seek
        }
      }

      nodes = append(nodes. n)
    }
    return nodes
}


func nextNode(it iterator.Iterator) *Node {
    for end := false; !end; end = !it.Next() {
      id, field := splitKey(it.Key())
      if field != nodeDBDiscoverRoot {
        continue
      }
      var n Node
      if err := rlp.DecodeBytes(it.Value(), &n); err != nil {
        log.Warn("Failed to decode node RLP", "id", id, "err", err)
        continue
      }
      return &n
    }
    return nil
  }


  

























