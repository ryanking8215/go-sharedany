package sharedany

import "sync"

type Pool[T any] struct {
	sync.Pool
	putNotify DoneFunc[T]
}

func NewPool[T any](factory FactoryFunc[T], putNotify DoneFunc[T]) *Pool[T] {
	p := Pool[T]{
		putNotify: putNotify,
	}
	p.Pool.New = func() interface{} {
		return New(factory(), 0, p.put)
	}
	return &p
}

func (p *Pool[T]) put(st *Shared[T]) {
	if p.putNotify != nil {
		p.putNotify(st)
	}
	p.Pool.Put(st)
}

func (p *Pool[T]) Get() *Shared[T] {
	return p.Pool.Get().(*Shared[T])
}
