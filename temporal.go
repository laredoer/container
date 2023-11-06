package container

import (
	"git.5th.im/lb-public/gear/log"

	"github.com/samber/do"
	"go.temporal.io/sdk/client"
)

const temporal = "Temporal"

func Workflow() client.Client {
	c, err := do.InvokeNamed[client.Client](container.injector, temporal)
	if err != nil {
		log.Panicf("Workflow invoke err:%v", err)
	}
	return c
}
