package servers

import (
	"framework/app/adapters/data"
	"framework/app/controllers"
	"framework/app/structs"
	"net/http"

	"github.com/1ets/lets"

	"github.com/gin-gonic/gin"
)

// HTTP Handler for get list of accounts
func HttpPostExample(g *gin.Context) {
	var httpRequest data.RequestExample
	var httpResponse data.ResponseExample
	var err error

	// Bind json body into struct format
	if err = g.Bind(&httpRequest); err != nil {
		lets.HttpResponseJson(g, httpResponse, err)
		return
	}

	var request structs.RequestIndex
	lets.Bind(httpRequest, &request)

	// Call example controller
	response, err := controllers.Index(request)
	if err != nil {
		httpResponse.Code = http.StatusBadRequest
		httpResponse.Message = err.Error()

		lets.HttpResponseJson(g, httpResponse, err)
		return
	}

	// Write json response
	lets.HttpResponseJson(g, response, err)
}
