package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {

	// Connect to a server
	nc, err := nats.Connect("ws://skidatahq.com:8081")
	if err != nil {
		log.Fatal("Cannot connect:", err)
	}

	// Simple Publisher
	nc.Publish("foo", []byte("Hello World from goclient!"))

	// // Simple Async Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	// // Responding to a request message
	// nc.Subscribe("request", func(m *nats.Msg) {
	// 	m.Respond([]byte("answer is 42"))
	// })

	// // Simple Sync Subscriber
	// sub, _ := nc.SubscribeSync("foo")
	// m, _ := sub.NextMsg(time.Second * 10)
	// fmt.Printf("%v\n", m)

	// // Channel Subscriber
	// ch := make(chan *nats.Msg, 64)
	// sub, _ = nc.ChanSubscribe("foo", ch)
	// msg := <-ch
	// fmt.Printf("%v\n", msg)

	// // Unsubscribe
	// sub.Unsubscribe()

	// // Drain
	// sub.Drain()

	// // Requests
	// msg, _ = nc.Request("help", []byte("help me"), 10*time.Millisecond)

	// // Replies
	// nc.Subscribe("help", func(m *nats.Msg) {
	// 	nc.Publish(m.Reply, []byte("I can help!"))
	// })

	// Drain connection (Preferred for responders)
	// Close() not needed if this is called.
	nc.Drain()

	// Close connection
	nc.Close()
}
