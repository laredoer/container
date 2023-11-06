package container

import (
	"git.5th.im/lb-public/gear/event"
	"git.5th.im/lb-public/gear/log"
	"github.com/samber/do"
)

const eventNode = "Event"

func Event() *event.Event {
	c, err := do.InvokeNamed[*event.Event](container.injector, eventNode)
	if err != nil {
		log.Panicf("Event invoke err:%v", err)
	}
	return c
}
