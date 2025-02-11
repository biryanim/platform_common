package validate

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ValidateErrors struct {
	Messages []string `json:"error_messages"`
}

func NewValidateErrors(messages ...string) *ValidateErrors {
	return &ValidateErrors{
		Messages: messages,
	}
}

func (v *ValidateErrors) Error() string {
	data, err := json.Marshal(v.Messages)
	if err != nil {
		return err.Error()
	}

	return string(data)
}

func IsValidateError(err error) bool {
	var ve *ValidateErrors
	return errors.As(err, &ve)
}

func (v *ValidateErrors) addError(message string) {
	v.Messages = append(v.Messages, message)
}
