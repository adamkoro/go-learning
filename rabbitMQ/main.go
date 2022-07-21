// code from https://dev.to/olushola_k/working-with-rabbitmq-in-golang-1kmj
package main

import amqp "github.com/streadway/amqp"

var url = "amqp://guest:guest@localhost:5672"

func main() {
	connection, err := amqp.Dial(url)
	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}

	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}

	// We publish the message to the exahange we created earlier
	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	// We create a queue named Test
	_, err = channel.QueueDeclare("test", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
		err = channel.QueueBind("test", "#", "events", false, nil)
	}

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
}
