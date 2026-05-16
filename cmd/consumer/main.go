package main

import "realtime-platform/internal/kafka"

func main() {

	consumer := kafka.NewConsumer()

	consumer.Start()
}