package transport

import (
	"github.com/streadway/amqp"
)

type Queue struct{
	AmqpConnection *amqp.Connection
	Channel *amqp.Channel
	Entities map[string]*amqp.Queue
}

func NewAmqpConnection(connstring string, queueNames []string)*Queue{
	conn, _ := amqp.Dial(connstring)
	ch, _ := conn.Channel()
	var (
		queue amqp.Queue
		queues = make(map[string]*amqp.Queue)
		)
	for _, channelName := range queueNames {
		queue , _ = ch.QueueDeclare(
			channelName,
			true,  /*durable connection*/
			false, /*auto delete*/
			false, /*exclusive*/
			false, /*nowai*t*/
			nil    /*arg amqp.table*/)
		queues[channelName] = &queue
	}

	return &Queue{AmqpConnection:conn, Channel:ch, Entities: queues}
}

func (q *Queue) SendMessageToQueue(message string, channelName string) error{
	publishedMsg := amqp.Publishing{
		DeliveryMode:2, // persistent msg delivery
		Body: []byte(message),
		Priority: 4,
		Type: "SendWelcomeEmail",
	}

	err := q.Channel.Publish(
		"",            // exchange string
		q.Entities[channelName].Name, // key string
		false,         // mandatory
		false,         // immediate
		publishedMsg)

	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) ConsumerMessageFromQueue(message string, queueName string) ([]interface{}, error){
	var response []interface{}

	msgs, err := q.Channel.Consume(
		q.Entities[queueName].Name, // queue string
		"", // consumer string
		true, // auto ack
		false, // exclusive bool
		false, // no local
		false, // no wait
		nil)

	if err != nil {
		return nil, err
	}

	for m := range msgs {
		response = append(response, m.Body)
	}

	return response, nil
}