package sharedany

import "sync/atomic"

const negativeRC = "sharedany: negative reference counter"

type FactoryFunc[T any] func() T
type DoneFunc[T any] func(st *Shared[T])

type Shared[T any] struct {
	d    T
	rc   int32
	done DoneFunc[T]
}

func New[T any](d T, rc int32, done DoneFunc[T]) *Shared[T] {
	if rc < 0 {
		panic(negativeRC)
	}
	return &Shared[T]{
		d:    d,
		rc:   rc,
		done: done,
	}
}

func (st *Shared[T]) Data() T {
	return st.d
}

func (st *Shared[T]) RC() int32 {
	return st.rc
}

func (st *Shared[T]) Add(delta int32) {
	if v := atomic.AddInt32(&st.rc, delta); v < 0 {
		panic(negativeRC)
	}
}

func (st *Shared[T]) Done() {
	v := atomic.AddInt32(&st.rc, -1)
	if v == 0 && st.done != nil {
		st.done(st)
	} else if v < 0 {
		panic(negativeRC)
	}
}
