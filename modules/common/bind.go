package common

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
)

func BindJsonError(err error, structName string) string {
	errorLog := fmt.Sprintf("%s", err)
	arr := strings.Split(errorLog, "\n")
	var result string
	var errorArr []string
	for _, v := range arr {
		index := strings.Index(v, structName)
		errorindex := strings.Index(v, "Error")
		arr := strings.Split(v[index:errorindex], ".")
		str := strings.Replace(arr[1], "'", "", 1)
		last := strcase.ToLowerCamel(str)
		errorArr = append(errorArr, last)
	}
	result = strings.Join(errorArr, ",")
	return result
}
