package sys

import (
	"github.com/biryanim/platform_common/pkg/sys/codes"
	"github.com/pkg/errors"
)

type commonError struct {
	msg  string
	code codes.Code
}

func NewCommonError(msg string, code codes.Code) *commonError {
	return &commonError{msg: msg, code: code}
}

func (r *commonError) Error() string {
	return r.msg
}

func (r *commonError) Code() codes.Code {
	return r.code
}

func IsCommonError(err error) bool {
	var ce *commonError
	return errors.As(err, &ce)
}

func GetCommonError(err error) *commonError {
	var ce *commonError
	if !errors.As(err, &ce) {
		return nil
	}

	return ce
}
