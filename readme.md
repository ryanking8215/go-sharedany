# Usage
```
go get github.com/ryanking8215/go-sharedany
```

# Shared anything
`Shared` is designed for the scenario that sharing anything to several consumers.

It is generic type with reference counter. When the reference counter is counting down to zero, the `Done` callback will be invoked.

If reference counter is negative, it panics.

A `byte slice` example
```golang
import "github.com/ryanking8215/sharedany"

data := make([]byte, 1024)
sharedBytes := sharedany.New[[]byte](data, 2, func(s *Shared[[]byte]) {
	// done callback
})

b := sharedBytes.Data()  // return data
rc := sharedBytes.RC() // return reference counter is 2

// after consumer1 consumes the data.
sharedBytes.Done()

// after consumer2 consumes the data.
sharedBytes.Done()

// sharedBytes's rc is down to zero, done callback will be invoked.
```

# Pool
`Pool` is a easy and efficient way to use `Shared`. `Shared` instance is given back to pool automaticlly when reference counter is counting down to zero.

A `byte slice` example again.
```
import "github.com/ryanking8215/sharedany"

create := func() []byte { return make([]byte, 1024) }
putNotify := func(sbs *sharedany.Shared[[]byte]) { println("put to pool")}

pool := sharedany.NewPool[[]byte](create, putNotify /* or nil if no need */)
sbs := pool.Get()

sbs.Add(2)

// after consumer1 consumes the buffer.
sbs.Done()

// after consumer2 consumes the buffer.
sbs.Done() 

// the shared instance is automaticlly given back to the pool.
// `putNotify` will be invoked

```