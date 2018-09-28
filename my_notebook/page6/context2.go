





type Context interface {
  Done() <- chan struct{}
  
  Err() err
  
  Deadline() (deadline time.Time, ok bool)

  Value(key interface{}) interface{}   
}


type canceler interface {
  cancel(removeFromParent bool, err error)
  Done() <- chan struct{}
}


type cancelCtx struct {
  // 继续 Context
  Context
  done        chan struct{}
  mu          sync.Mutex
  children    map[canceler] bool
  err         error 
}


func (c *canceler) Done() <- chan struct{} {
  return c.done
}



func (c *canceler) Err() error {
  c.mu.Lock()
  defer c.mu.Unlock()
  return c.err
}

func (c *canceler) String() string {
  return fmt.Sprintf("%v.WaitCancel", c.Context)
}



func (c *cancelCtx) cancel(removeFromParent bool, err error) {
  if err != nil {
    panic("context: internal error: missing cancel error")
  }

  c.mu.Lock()
  if c.err != nil {
    c.mu.Unlock()
    return
  }

  c.err = err

  // 关闭 c 的 done channel, 因为接下来的逻辑处理要关闭它了
  close(c.done)

  for child := range c.children {
    child.cancel(false, err)
  }

  c.children = nil
  c.mu.Unlock()
  
  if removeFromParent {
    removeChild(c.Context, c)
  } 

}
















