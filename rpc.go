package container

import (
	"git.5th.im/lb-public/gear/log"
	client "github.com/micro/go-micro/client"
	"github.com/samber/do"
)

type NewServiceFn[T any] func(name string, c client.Client) T

func (f NewServiceFn[T]) Resolve(name string) T {
	return f(name, container.client)
}

func ProvideRPC[T any](f NewServiceFn[T], names ...string) {
	var name string
	if len(names) != 0 {
		name = names[0]
	}
	do.OverrideValue(container.injector, f.Resolve(name))
}

// Invoke  调用 rpc
func Invoke[T any]() T {
	var v T
	v, err := do.Invoke[T](container.injector)
	if err != nil {
		log.Panicf("rpc invoke err:%v", err)
	}
	return v
}
