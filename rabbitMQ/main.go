// code from https://dev.to/olushola_k/working-with-rabbitmq-in-golang-1kmj
package main

import (
	"fmt"

	amqp "github.com/streadway/amqp"
)

var url = "amqp://guest:guest@localhost:5672"

func main() {
	fmt.Println("Create connection")
	connection, err := amqp.Dial(url)

	if err != nil {
		panic("could not establish connection with RabbitMQ:" + err.Error())
	}

	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the
	// connection itself
	fmt.Println("Create channel")
	channel, err := connection.Channel()

	if err != nil {
		panic("could not open RabbitMQ channel:" + err.Error())
	}
	fmt.Println("Declare channel")
	// We create an exahange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("events", "topic", true, false, false, false, nil)

	if err != nil {
		panic(err)
	}

	// We create a message to be sent to the queue.
	// It has to be an instance of the aqmp publishing struct
	fmt.Println("Publish message")
	message := amqp.Publishing{
		Body: []byte("Hello World"),
	}

	// We publish the message to the exahange we created earlier
	fmt.Println("Publish channel")
	err = channel.Publish("events", "random-key", false, false, message)

	if err != nil {
		panic("error publishing a message to the queue:" + err.Error())
	}

	// We create a queue named Test
	fmt.Println("Create queue")
	_, err = channel.QueueDeclare("test", true, false, false, false, nil)

	if err != nil {
		panic("error declaring the queue: " + err.Error())
	}
	fmt.Println("Bind queue to exchange")
	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("test", "#", "events", false, nil)

	if err != nil {
		panic("error binding to the queue: " + err.Error())
	}
	fmt.Println("Get data from queue")
	// We consume data from the queue named Test using the channel we created in go.
	msgs, err := channel.Consume("test", "", false, false, false, false, nil)

	if err != nil {
		panic("error consuming the queue: " + err.Error())
	}
	fmt.Println("Show messages data")
	// We loop through the messages in the queue and print them in the console.
	// The msgs will be a go channel, not an amqp channel
	for msg := range msgs {
		fmt.Println("message received: " + string(msg.Body))
		msg.Ack(true)
	}

	// We close the connection after the operation has completed.
	defer connection.Close()
}
