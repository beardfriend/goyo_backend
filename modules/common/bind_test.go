package common

import (
	"errors"
	"testing"
)

func TestBindMessage(t *testing.T) {
	errorMessage := `Key: 'RegistRequestBody.Day' Error:Field validation for 'Day' failed on the 'required' tag
	Key: 'RegistRequestBody.YogaName' Error:Field validation for 'YogaName' failed on the 'required' tag
	Key: 'RegistRequestBody.StartTime' Error:Field validation for 'StartTime' failed on the 'required' tag
	Key: 'RegistRequestBody.EndTime' Error:Field validation for 'EndTime' failed on the 'required' tag`
	error := errors.New(errorMessage)
	dd := BindJsonError(error, "RegistRequestBody")
	t.Logf("\n-----------------\n%s\n-----------------", dd)
}
