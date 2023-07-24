package main

import (
	"log"

	signals "github.com/Arif9878/temporal-signal"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// The client and worker are heavyweight objects that should be created once per process.
	c, err := client.Dial(client.Options{
		HostPort:  "localhost:7233",
		Namespace: "user-namespace",
	})
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}
	defer c.Close()

	w := worker.New(c, "await_signals", worker.Options{})

	w.RegisterWorkflow(signals.AwaitSignalsWorkflow)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
