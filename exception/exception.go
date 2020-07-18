package exception

import (
	"io"
)

type Exception interface {
}

type trier struct {
	err Exception
}

type tryResource struct {
	trier
	closer io.Closer
}

// for try...catch...finally as java
// support chain call
func NewTrier() *trier {
	return new(trier)
}

// for try...with...resource as java
// support chain call
func NewTryResource(closer io.Closer) *tryResource {
	return &tryResource{
		closer: closer,
	}
}

func (t *trier) Try(f func()) (trier *trier) {
	defer func() {
		trier = t
		if err := recover(); err != nil {
			t.err = err
		}
	}()

	f()
	return t
}

func (t *trier) Throw(e Exception) {
	if e != nil {
		panic(e)
	}
}

func (t *trier) Catch(f func(e Exception)) *trier {
	if t.err != nil {
		f(t.err)
	}
	return t
}

func (t *trier) Finally(f func()) {
	f()
}

func (t *tryResource) Try(f func()) (trs *tryResource) {
	defer func() {
		trs = t
		if err := recover(); err != nil {
			t.err = err
		}
		if t.closer != nil {
			err := t.closer.Close()
			if err != nil && t.err == nil {
				t.err = err
			}
		}
	}()

	f()
	return t
}
