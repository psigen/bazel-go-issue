package main

import (
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

func main() {
	// This is just some boilerplate to minimally use the library.
	c, err := client.NewClient(client.Options{})
	defer c.Close()

	w := worker.New(c, "hello-world", worker.Options{})
	w.Run(worker.InterruptCh())
}
