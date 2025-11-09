package sync

import "sync"

type Pool[T any] struct {
	p sync.Pool
}

func (p *Pool[T]) Get() T {
	v := p.p.Get()
	if v == nil {
		return *new(T)
	}
	return v.(T)
}

func (p *Pool[T]) Put(x T) {
	p.p.Put(x)
}
