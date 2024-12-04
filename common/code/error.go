package code

import (
	"encoding/json"
	"fmt"
)

type AppError struct {
	code  int
	msg   string
	args  []any
	cause error
}

func (e *AppError) Code() int {
	return e.code
}

func (e *AppError) Msg() string {
	if len(e.args) != 0 {
		return fmt.Sprintf(e.msg, e.args...)
	}
	return e.msg
}

func (e *AppError) String() string {
	return e.Error()
}

func (e *AppError) Error() string {
	formatErr := struct {
		Code  int    `json:"code"`
		Msg   string `json:"msg"`
		Cause string `json:"cause"`
	}{
		e.code, e.msg, "",
	}

	if e.cause != nil {
		formatErr.Cause = e.cause.Error()
	}

	bs, _ := json.Marshal(formatErr)
	return string(bs)
}

func newError(code int, msg string) *AppError {
	if _, ok := codeMap[code]; ok {
		panic("重复的错误码!")
	}

	appErr := &AppError{
		code:  code,
		msg:   msg,
		cause: nil,
	}
	codeMap[code] = appErr
	return appErr
}

func (e *AppError) WithCause(err error) *AppError {
	newErr := e.Clone()
	newErr.cause = err
	return newErr
}

func (e *AppError) WithArgs(args ...any) *AppError {
	e.args = args
	return e
}

func (e *AppError) Clone() *AppError {
	return &AppError{code: e.code, msg: e.msg, cause: e.cause}
}

// Wrap 包装下非自定义调用错误
func Wrap(msg string, err error) *AppError {
	return &AppError{
		code:  -1,
		msg:   msg,
		cause: err,
	}
}

func (e *AppError) Equal(other *AppError) bool {
	return e.code == other.code
}
