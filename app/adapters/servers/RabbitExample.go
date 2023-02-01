package servers

import (
	"fmt"
	"framework/app/adapters/clients"
	"framework/app/adapters/data"
	"framework/app/controllers"

	"github.com/1ets/lets"
	"github.com/1ets/lets/types"
)

// Adapter for RabbitMQ consumer.
func RabbitExample(r *types.Event) {
	var request data.RequestExample
	var err error

	lets.Bind(r.GetData(), &request)

	// Call controller
	response, err := controllers.RabbitConsumerExample(request)
	if err != nil {
		lets.LogE("gRPC Server: GrpcExample.Greeting:", err.Error())
		return
	}

	if err != nil {
		fmt.Println("ERR: ", err.Error())
	}

	// Create reply for sync
	if r.GetReplyTo() != nil {
		clients.RabbitExample.GreetingCallback(r, &response)
	}
}
