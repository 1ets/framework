package servers

import (
	"framework/app/adapters/clients"
	"framework/app/adapters/data"
	"framework/app/controllers"

	"github.com/1ets/lets"
	"github.com/1ets/lets/types"
)

// Adapter for RabbitMQ consumer "transfer-request".
func RabbitSagaExampleTransfer(r *types.Event) {
	var request data.RequestTransfer
	var err error

	lets.Bind(r.GetData(), &request)

	// Call controller
	response, err := controllers.RabbitSagaTransferServiceProvider(request)
	if err != nil {
		lets.LogE("RabbitMQ Server: RabbitSagaExampleTransfer:", err.Error())
		return
	}

	// Create reply
	if r.GetReplyTo() != nil {
		err := clients.RabbitSagaExampleCallback.TransactionCallback(r, &response)
		if err != nil {
			lets.LogE("ReplyError: ", err.Error())
		}
	}
}

// Adapter for RabbitMQ consumer "balance-request".
func RabbitSagaExampleBalance(r *types.Event) {
	var request data.RequestBalance
	var err error

	lets.Bind(r.GetData(), &request)

	// Call controller
	response, err := controllers.RabbitSagaBalanceServiceProvider(request)
	if err != nil {
		lets.LogE("RabbitMQ Server: RabbitBalanceRequest:", err.Error())
		return
	}

	// Create reply
	if r.GetReplyTo() != nil {
		err := clients.RabbitSagaExampleCallback.BalanceCallback(r, &response)
		if err != nil {
			lets.LogE("ReplyError: ", err.Error())
		}
	}
}

// Adapter for RabbitMQ consumer "balance-request".
func RabbitSagaExampleNotification(r *types.Event) {
	var request data.RequestNotification
	var err error

	lets.Bind(r.GetData(), &request)

	// Call controller
	response, err := controllers.RabbitSagaNotificationServiceProvider(request)
	if err != nil {
		lets.LogE("RabbitMQ Server: RabbitNotificationRequest:", err.Error())
		return
	}

	// Create reply
	if r.GetReplyTo() != nil {
		err := clients.RabbitSagaExampleCallback.NotificationCallback(r, &response)
		if err != nil {
			lets.LogE("ReplyError: ", err.Error())
		}
	}
}
