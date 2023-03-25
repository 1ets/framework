package controllers

import (
	"fmt"
	"framework/app/structs"

	"github.com/pkg/errors"
)

// Example controller.
func Index(request structs.RequestIndex) (response structs.ResponseIndex, err error) {
	if name := request.Name; request.Name != "" {
		response.Greeting = fmt.Sprintf("Hello %s, how are you!", name)
		return
	}

	err = errors.New("Name Cannot be empty")

	return
}
