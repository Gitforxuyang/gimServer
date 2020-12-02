package rabbitmq

import (
	"fmt"
	"gimServer/conf"
	"gimServer/infra/utils"
	"github.com/streadway/amqp"
)

type Queue struct {
	ch *amqp.Channel
}

func (m *Queue) Publish(routingKey string, header amqp.Table, body []byte) error {
	fmt.Println(routingKey)
	fmt.Println(header)
	return m.ch.Publish("im.topic", routingKey, false, false, amqp.Publishing{
		DeliveryMode: 2,
		Headers:      header,
		Body:         body,
	})
}
func InitClient(config *conf.Config) *Queue {
	c, err := amqp.Dial(config.Rabbit.Addr)
	utils.Must(err)
	ch, err := c.Channel()
	utils.Must(err)
	err = ch.ExchangeDeclare("im.topic", "topic", true, false, false, false, nil)
	utils.Must(err)
	mq := Queue{}
	mq.ch = ch
	return &mq
}
