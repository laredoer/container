package container

import (
	"git.5th.im/lb-public/gear/log"
	"github.com/robfig/cron/v3"
	"github.com/samber/do"
)

const cronNode = "CronNode-"

type CronNode interface {
	GetCronNode() string
}

func CronFunc[C CronNode](spec string, cmd func()) error {
	var cn C
	c, err := do.InvokeNamed[*cron.Cron](container.injector, cronNode+cn.GetCronNode())
	if err != nil {
		log.Panicf("cron invoke err:%v", err)
	}
	if _, err = c.AddFunc(spec, cmd); err != nil {
		return err
	}
	return nil
}

func StartCron[C CronNode]() {
	var cn C
	c, err := do.InvokeNamed[*cron.Cron](container.injector, cronNode+cn.GetCronNode())
	if err != nil {
		log.Panicf("cron invoke err:%v", err)
	}
	c.Start()
}

func StopCron[C CronNode]() {
	var cn C
	c, _ := do.InvokeNamed[*cron.Cron](container.injector, cronNode+cn.GetCronNode())
	if c == nil {
		return
	}
	c.Stop()
}
