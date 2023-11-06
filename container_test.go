package container

import (
	"context"
	"fmt"
	"testing"
	"time"

	"git.5th.im/lb-public/gear/log"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/robfig/cron/v3"
	. "github.com/smartystreets/goconvey/convey"
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

type LarkActivity struct{}

func (l LarkActivity) GetLarkNode() string {
	return "activity"
}

type CronReward struct{}

func (c CronReward) GetCronNode() string {
	return "reward"
}

func TestContainer(t *testing.T) {

	Convey("TestContainer", t, func() {

		Convey("TestGetDB", func() {
			mockRewardDB, _, _ := sqlmock.New()
			New(WithTestDB[Reward](mockRewardDB))
			Convey("GetDB result is not nil", func() {
				db := GetDB[Reward](context.Background())
				So(db, ShouldNotBeNil)
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
				WithLark[LarkReward]("https://open.feishu.cn/open-apis/bot/v2/hook/2714a5fb-07a4-4ca8-8b72-265cddec5385"),
				WithLark[LarkActivity]("https://open.feishu.cn/open-apis/bot/v2/hook/b42f3768-2695-4a9c-978e-abbd151dc1fc"),
			)
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

	})

}
