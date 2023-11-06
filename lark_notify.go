package container

import (
	"fmt"

	"git.5th.im/lb-public/gear"
	"git.5th.im/lb-public/gear/feishu"
	"git.5th.im/lb-public/gear/log"
	"github.com/samber/do"
)

const larkNode = "LarkNode-"

type LarkNode interface {
	GetLarkNode() string
}

func LarkNotify[L LarkNode](title, message string, at ...string) {
	if !gear.Env.IsProd() {
		return
	}

	var l L
	alarms, err := do.InvokeNamed[*feishu.Client](container.injector, larkNode+l.GetLarkNode())
	if err != nil {
		log.Errorf("Get Lark alarms err:%v", err)
		return
	}

	if err := alarms.Notify(fmt.Sprintf("[%s] %s", gear.Env.String(), title), message, at...); err != nil {
		log.Errorf("发送飞书通知失败 err:%v", err)
	}
}
