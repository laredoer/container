package container

import (
	"context"
	"fmt"
	"testing"
	"time"

	"git.5th.im/lb-public/gear/log"
	"git.5th.im/lb-public/gear/mq/rabbitconsumer"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/robfig/cron/v3"
	. "github.com/smartystreets/goconvey/convey"
	"go.temporal.io/sdk/client"
)

type Reward struct{}

func (r Reward) GetDBName() string {
	return "reward"
}

func (r Reward) GetDialect() string {
	return "postgres"
}

type LarkReward struct{}

func (l LarkReward) GetLarkNode() string {
	return "reward"
}

func (l LarkReward) HookAddr() string {
	return "https://open.feishu.cn/open-apis/message/v4/send/"
}

type LarkActivity struct{}

func (l LarkActivity) GetLarkNode() string {
	return "activity"
}

func (l LarkActivity) HookAddr() string {
	return "https://open.feishu.cn/open-apis/message/v4/send/"
}

type CronReward struct{}

func (c CronReward) GetCronNode() string {
	return "reward"
}

type ActivityAgg struct {
	ID   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Num  string `json:"num"`
}

type ActivityRewardCount struct {
	Env  string         `json:"env"`
	Date string         `json:"date"`
	List []*ActivityAgg `json:"list"`
}

func (ActivityRewardCount) TemplateID() string {
	return "ctp_AA85XZHdHbcz"
}

func TestContainer(t *testing.T) {

	Convey("TestContainer", t, func() {

		Convey("TestGetDB", func() {
			New(WithTestDB[Reward]())
			Convey("GetDB result is not nil", func() {
				db := GetDB[Reward](context.Background())
				So(db, ShouldNotBeNil)

				sqlmock := SQLMock[Reward]()
				So(sqlmock, ShouldNotBeNil)

				RefreshSQLMock[Reward]()

				sqlmock2 := SQLMock[Reward]()
				So(sqlmock2, ShouldNotBeNil)

				So(sqlmock, ShouldNotEqual, sqlmock2)

			})
		})

		Convey("Test Redis", func() {
			New(
				WithTestRedis("lb"),
			)
			Convey("GetRedis result is not nil", func() {
				rcli := GetRedis(context.Background())
				So(rcli, ShouldNotBeNil)
			})
			Convey("GetCache result is not nil", func() {
				cache := GetCache()
				So(cache, ShouldNotBeNil)
			})

			Convey("Test Fetch nomal", func() {
				res, err := Fetch(context.Background(), "key", time.Minute, func() (rawResult Reward, err error) {
					return Reward{}, nil
				})
				Convey("error should be nil", func() {
					So(err, ShouldBeNil)
				})
				Convey("res should not be nil", func() {
					So(res, ShouldNotBeNil)
				})
			})

			Convey("Test Fetch has error", func() {
				res, err := Fetch(context.Background(), "key", time.Minute, func() (rawResult Reward, err error) {
					return Reward{}, fmt.Errorf("system error")
				})
				Convey("err should be error", func() {
					So(err, ShouldBeError)
				})
				Convey("res should be nil", func() {
					So(res, ShouldBeNil)
				})
			})
		})

		Convey("Test Lark", func() {
			New(
				WithLark[LarkActivity](),
			)

			CardNotify[LarkActivity](ActivityRewardCount{
				Env:  "test",
				Date: "2023-11-08",
				List: []*ActivityAgg{
					{
						ID:   "1",
						Code: "code1",
						Name: "name1",
						Num:  "1",
					},
					{
						ID:   "1",
						Code: "code2",
						Name: "name2",
						Num:  "2",
					},
				},
			})
		})

		Convey("Test Cron", func() {
			New(
				WithCron[CronReward](cron.WithSeconds()),
			)

			Convey("Cron can run", func() {
				CronFunc[CronReward]("@every 1s", func() {
					log.Info("cron is runing")
				})
				StartCron[CronReward]()
				time.Sleep(time.Second * 2)
			})
		})

		Convey("Test MQ", func() {
			New(
				WithTestMQ(),
			)

			Convey("Register", func() {
				RegisterConsumer("AccountOpen", func(body []byte, headers map[string]interface{}) error {
					log.Info("AccountOpen", string(body))
					return nil
				}, rabbitconsumer.WithPreHooks(func(ctx context.Context, msg amqp.Delivery) (requeue bool, err error) {
					msg.Headers["body"] = string(msg.Body)
					log.Infof("AccountOpen pre hook %s", msg.Body)
					return
				}))
			})

			Convey("Push", func() {
				So(func() {
					SendQueueMessage(context.Background(), "test-push", nil, ActivityAgg{})
				}, ShouldNotPanic)

			})
		})

		Convey("Test LocalCache", func() {
			New(
				WithLocalCache(),
			)

			Convey("LocalCache set get", func() {
				Convey("String set get", func() {
					SetLocalCache("key1", "value1", time.Second*3)
					v, ok := GetLocalCache[string]("key1")
					So(ok, ShouldBeTrue)
					So(v, ShouldEqual, "value1")
				})
				Convey("Struct set get", func() {
					SetLocalCache("key3", ActivityAgg{
						ID:   "1",
						Code: "code",
						Name: "name",
						Num:  "1",
					}, time.Second*3)
					v3, ok := GetLocalCache[ActivityAgg]("key3")
					So(ok, ShouldBeTrue)
					So(v3, ShouldEqual, ActivityAgg{
						ID:   "1",
						Code: "code",
						Name: "name",
						Num:  "1",
					})
				})
			})
		})

		Convey("Test temporal", func() {
			New(WithTestTemporal())
			Convey("Test temporal workflow", func() {
				So(func() {
					Workflow().ExecuteWorkflow(context.Background(),
						client.StartWorkflowOptions{
							TaskQueue: "Test_Task_Queue",
						},
						func() {},
					)
				}, ShouldNotPanic)
			})
		})

	})
}
