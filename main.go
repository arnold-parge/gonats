package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	// Define request handler
	echoHandler := func(req micro.Request) {
		log.Println("Received request:", string(req.Data()))
		req.Respond(req.Data())
	}

	// Create a service
	service, err := micro.AddService(nc, micro.Config{
		Name:    "EchoService",
		Version: "1.0.0",
		Endpoint: &micro.EndpointConfig{
			Subject: "svc.echo",
			Handler: micro.HandlerFunc(echoHandler),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Service running:", service.Info().Name)
	select {} // Keep the service running
}
