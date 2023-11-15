package container

import (
	"database/sql"

	"time"

	"git.5th.im/lb-public/gear"
	"git.5th.im/lb-public/gear/cache"
	db "git.5th.im/lb-public/gear/db/v2"
	"git.5th.im/lb-public/gear/event"
	"git.5th.im/lb-public/gear/feishu"
	"git.5th.im/lb-public/gear/log"
	"git.5th.im/lb-public/gear/mq"
	"git.5th.im/lb-public/gear/mq/rabbitmq"
	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis/v8"
	client "github.com/micro/go-micro/client"
	"github.com/robfig/cron/v3"
	"github.com/samber/do"
	"github.com/sony/sonyflake"
	workflowClient "go.temporal.io/sdk/client"
	"golang.org/x/exp/constraints"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var container *_Container

// 服务容器
type _Container struct {
	injector *do.Injector
	client   client.Client
	*sonyflake.Sonyflake
}

func New(ops ...Op) {
	container = &_Container{
		injector:  do.New(),
		Sonyflake: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}

	for _, op := range ops {
		op(container)
	}
}

type Op func(c *_Container)

func WithDB[T DB]() Op {
	var t T
	resourceName := t.GetDBName()
	dialect := t.GetDialect()
	return func(c *_Container) {
		switch dialect {
		case "mysql":
			do.OverrideNamedValue(c.injector, resourceName, db.NewMySQLFromConfig(resourceName))
		case "postgres":
			do.OverrideNamedValue(c.injector, resourceName, db.NewPGFromConfig(resourceName))
		}
	}
}

func WithTestDB[T DB](d *sql.DB) Op {
	var t T
	resourceName := t.GetDBName()
	dialect := t.GetDialect()
	return func(c *_Container) {
		switch dialect {
		case "mysql":
			do.OverrideNamedValue(c.injector, resourceName, &db.DB{DB: getMockGorm(d)})
		case "postgres":
			do.OverrideNamedValue(c.injector, resourceName, &db.DB{DB: getMockGorm(d)})
		}
	}
}

func getMockGorm(mockDB *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		Conn:       mockDB,
		DriverName: "postgres",
	})
	gormDB, _ := gorm.Open(dialector, &gorm.Config{})
	return gormDB
}

func WithRedis(prefix, resourceName string) Op {
	return func(c *_Container) {
		redis := cache.NewRedisFromConfig(resourceName)
		do.OverrideNamedValue(c.injector, redisNode, redis)
		do.OverrideNamedValue(c.injector, cacheNode,
			cache.NewCache(redis.Client, &cache.Config{
				Prefix:     prefix,
				DefaultTTL: 24 * time.Hour,
			}))
	}
}

func WithTestRedis(prefix string) Op {
	return func(c *_Container) {
		mr, err := miniredis.Run()
		if err != nil {
			panic(err)
		}
		redisClient := redis.NewClient(&redis.Options{
			Addr: mr.Addr(),
		})
		redis := &cache.RedisClient{Client: redisClient}
		do.OverrideNamedValue(c.injector, redisNode, redis)
		do.OverrideNamedValue(c.injector, cacheNode,
			cache.NewCache(redis.Client, &cache.Config{
				Prefix:     prefix,
				DefaultTTL: 24 * time.Hour,
			}))
	}
}

func WithClient(client client.Client) Op {
	return func(c *_Container) {
		c.client = client
	}
}

func WithTemporal(hostPort string) Op {
	return func(c *_Container) {
		clientOptions := workflowClient.Options{
			HostPort: hostPort,
		}
		if gear.Env.IsDev() {
			clientOptions.HostPort = "127.0.0.1:7233"
		}

		wcli, err := workflowClient.Dial(clientOptions)
		if err != nil {
			log.Fatalf("failed new workflow client err: %s", err)
		}

		do.OverrideNamedValue(c.injector, temporal, wcli)
	}
}

func WithMQ(resourceName string) Op {
	return func(c *_Container) {
		mqCli = &mqClient{
			rabbitmq:    rabbitmq.NewClientFromConfig(resourceName),
			producerMap: make(map[string]*producerCache),
			consumerMap: make(map[string]mq.Consumer),
		}
		go mqCli.checkProducerCache()
	}
}

func WithTestMQ() Op {
	return func(c *_Container) {
		mqCli = &mqClient{
			producerMap: make(map[string]*producerCache),
			consumerMap: make(map[string]mq.Consumer),
			isTestMode:  true,
		}
		go mqCli.checkProducerCache()
	}
}

func WithLark[L LarkNode]() Op {
	return func(c *_Container) {
		var l L
		do.OverrideNamedValue(c.injector, larkNode+l.GetLarkNode(), feishu.New(l.HookAddr()))
	}
}

func WithCron[C CronNode](opts ...cron.Option) Op {
	return func(c *_Container) {
		var cn C
		do.OverrideNamedValue(c.injector, cronNode+cn.GetCronNode(), cron.New(opts...))
	}
}

func WithEvent(rn string) Op {

	evt, err := event.NewEventFromConfig(rn)
	if err != nil {
		log.Panic(err)
	}

	return func(c *_Container) {
		do.OverrideNamedValue(c.injector, eventNode, evt)
	}
}

// Close 关闭连接
func Close() {
	if mqCli != nil {
		mqCli.close()
	}
}

func NextID[T constraints.Integer]() T {
	id, _ := container.NextID()
	return T(id)
}
