// context包

type Context interface {
  // 返回已经取消了的
  Done() <- chan struct{}

  Err() error
}