package utils

import (
	"fmt"

	"github.com/photo_catalog/pkg/v1/responses"
)

func NewErrorMessage(name, desc string, errors interface{}) responses.ErrorMessage {
	var errorList []string

	if errList, ok := errors.([]string); ok {
		errorList = errList
	} else {
		errorList = []string{fmt.Sprintf("%v", errors)}
	}
	return responses.ErrorMessage{
		Name:  name,
		Desc:  desc,
		Error: errorList,
	}
}
