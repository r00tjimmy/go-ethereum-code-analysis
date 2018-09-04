// context包

type Context interface {
  // 返回已经取消了的
  Done() <- chan struct{}

  Err() error

  Deadline() (deadline time.Time, ok bool)

  Value(key interface{}) interface{}
}


func WithCancel(parent Context) (ctx Context, cancel CancelFunc) {
  c := newCancelCtx(parent)
  propagateCancel(parent, &c)
  return &c, func() { c.cancel(true, Canceled) }
}


func newCancelCtx(parent Context) cancelCtx {
  return cancelCtx {
    Context:    parent,
    done:       make(chan struct{}), 
  }
}


