package container

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"git.5th.im/lb-public/gear"
	"git.5th.im/lb-public/gear/feishu"
	"git.5th.im/lb-public/gear/log"
	jsoniter "github.com/json-iterator/go"
	"github.com/samber/do"
)

const larkNode = "LarkNode-"

type LarkNode interface {
	GetLarkNode() string
	HookAddr() string
}

func Notify[L LarkNode](title, message string, at ...string) {
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

type Template interface {
	TemplateID() string
}

func CardNotify[L LarkNode, T Template](template T) {
	data := map[string]any{
		"msg_type": "interactive",
		"card": map[string]any{
			"type": "template",
			"data": map[string]any{
				"template_id":       template.TemplateID(),
				"template_variable": template,
			},
		},
	}

	var l L
	jsonMsg, _ := jsoniter.Marshal(data)
	req, err := http.NewRequest("POST", l.HookAddr(), strings.NewReader(string(jsonMsg)))
	if err != nil {
		log.Errorf("send feishu message %v: %v", jsonMsg, err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	httpcli := &http.Client{Timeout: 5 * time.Second}
	resp, err := httpcli.Do(req)
	if err != nil {
		log.Errorf("send feishu message %v: %v", jsonMsg, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("request feishu webhook fail: %s, message: %s", resp.Status, jsonMsg)
	}
}
