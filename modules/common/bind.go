package common

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func BindJsonError(err error, structName string) string {
	errorMessage := fmt.Sprintf("%s", err)

	splitMessage := strings.Split(errorMessage, "\n")

	var errorMessageArray []string

	for _, v := range splitMessage {
		indexStructNameStart := strings.Index(v, structName)
		indexErrorStart := strings.Index(v, "Error")

		commaSplit := strings.Split(v[indexStructNameStart:indexErrorStart], ".")
		message := strings.Replace(commaSplit[1], "'", "", 1)
		messageValue := strcase.ToLowerCamel(message)

		errorMessageArray = append(errorMessageArray, messageValue)
	}
	return strings.Join(errorMessageArray, ",")
}
