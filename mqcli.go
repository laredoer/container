// Package mqcli is a wrapper of gear/mq
package container

import (
	"context"

	"sync"
	"time"

	"git.5th.im/lb-public/gear/mq/rabbitmq"
	jsoniter "github.com/json-iterator/go"

	"git.5th.im/lb-public/gear/log"
	"git.5th.im/lb-public/gear/mq"
	"git.5th.im/lb-public/gear/mq/rabbitcli"

	"git.5th.im/lb-public/gear/mq/rabbitconsumer"
)

var mqCli *mqClient

type mqClient struct {
	sync.RWMutex
	rabbitmq    *rabbitcli.RabbitMQClient
	producerMap map[string]*producerCache
	consumerMap map[string]mq.Consumer
}

type producerCache struct {
	expireTime int64
	producer   mq.Producer
}

const expireDuration = 10 * 60 // 缓存 10m

// 5m ticker
func (m *mqClient) checkProducerCache() {
	c := time.Tick(time.Minute * 5)
	for next := range c {
		now := next.Unix()
		for k, v := range m.producerMap {
			if v.expireTime >= now {
				continue
			}
			m.Lock()
			if v.expireTime < now { // 二次判断
				delete(m.producerMap, k)
			}
			m.Unlock()
			v.producer.Close() // unlock 之后 close
			log.Info("close producer: ", k)
		}
	}
}

// 从缓存获取 producer
func (m *mqClient) getProducer(resourceName string) mq.Producer {
	m.Lock()
	defer m.Unlock()
	pcache, has := m.producerMap[resourceName]
	if !has {
		pcache = &producerCache{
			producer: rabbitmq.NewProducerFromConfig(m.rabbitmq, resourceName),
		}
		m.producerMap[resourceName] = pcache
	}
	pcache.expireTime = time.Now().Unix() + expireDuration
	return pcache.producer
}

// 消费 message
func (m *mqClient) newContextConsumer(resourceName string, processor rabbitconsumer.ContextProcessor) {
	consumer := rabbitmq.NewConsumerFromConfig(m.rabbitmq, resourceName, processor)
	consumer.Start()
	m.consumerMap[resourceName] = consumer
}

func (m *mqClient) newConsumer(resourceName string, processor rabbitconsumer.Processor) {
	consumer := rabbitmq.NewConsumerFromConfig(m.rabbitmq, resourceName, processor)
	consumer.Start()
	m.consumerMap[resourceName] = consumer
}

type Processor interface {
	rabbitconsumer.ContextProcessor | rabbitconsumer.Processor |
		func(ctx context.Context, msg []byte, headers map[string]interface{}) error |
		func(msg []byte, headers map[string]interface{}) error
}

// RegisterConsumer 注册消费者
func RegisterConsumer[F Processor](mqResourceName string, processor F) {
	if mqCli == nil {
		log.Panicf("MQClient Not Init mqResourceName:%s", mqResourceName)
	}

	switch f := any(processor).(type) {
	case rabbitconsumer.ContextProcessor:
		mqCli.newContextConsumer(mqResourceName, f)
	case func(ctx context.Context, msg []byte, headers map[string]interface{}) error:
		mqCli.newContextConsumer(mqResourceName, f)
	case rabbitconsumer.Processor:
		mqCli.newConsumer(mqResourceName, f)
	case func(msg []byte, headers map[string]interface{}) error:
		mqCli.newConsumer(mqResourceName, f)
	}
}

// SendQueueMessage 发送消息到队列
func SendQueueMessage[T any](ctx context.Context, resourceName string, headers map[string]any, values ...T) {

	if mqCli == nil {
		log.Panicf("MQClient Not Init resourceName:%s", resourceName)
	}

	for _, v := range values {
		body, _ := jsoniter.Marshal(v)
		if err := mqCli.rabbitMQPublishWithHeaders(ctx, resourceName, headers, body); err != nil {
			log.Errorf("mq %s publish: %s err:%v", resourceName, body, err)
		}
	}

}

// 发布 message，带有 headers
func (m *mqClient) rabbitMQPublishWithHeaders(ctx context.Context, resourceName string, headers map[string]interface{}, body []byte) error {
	producer := m.getProducer(resourceName)
	err := producer.Publish(ctx, &mq.Message{
		Headers: headers,
		Body:    body,
	})
	log.Infof("mq %s publish: %s", resourceName, body)
	return err
}

// Close 关闭连接
func (m *mqClient) close() {
	for _, v := range m.consumerMap {
		v.Close()
	}
	for _, v := range m.producerMap {
		v.producer.Close()
	}
}

func GetMQCli() *rabbitcli.RabbitMQClient {
	return mqCli.rabbitmq
}
