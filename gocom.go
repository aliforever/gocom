package gocom

import (
	"github.com/aliforever/gocom/keyval"
	"github.com/aliforever/gocom/pubsub"
	"github.com/aliforever/gocom/queue"
)

func KeyVal(name ...string) keyval.KeyValClient {
	return keyval.Get(name...)
}

func PubSub(name ...string) pubsub.PubSubClient {
	return pubsub.Get(name...)
}

func Queue(name ...string) queue.QueueClient {
	return queue.Get(name...)
}
