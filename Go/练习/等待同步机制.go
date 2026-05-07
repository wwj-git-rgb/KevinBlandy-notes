type Future[T any] struct {
	val  atomic.Pointer[T]
	done chan struct{}
}

func NewFuture[T any]() *Future[T] {
	return &Future[T]{done: make(chan struct{})}
}

func (f *Future[T]) Set(val *T) bool {
	if val != nil && f.val.CompareAndSwap(nil, val) {
		close(f.done)
		return true
	}
	return false
}

func (f *Future[T]) Get(ctx context.Context) *T {
	if v := f.val.Load(); v != nil {
		return v
	}
	select {
	case <-f.done:
	case <-ctx.Done():
	}
	return f.val.Load()
}

func (f *Future[T]) TryGet() (*T, bool) {
	v := f.val.Load()
	return v, v != nil
}
