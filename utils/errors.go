package utils

import "fmt"

// XError the common error
type XError struct {
	errCode uint32
	errMsg  string
}

func (xe *XError) Error() string {
	return fmt.Sprintf("error code: %d error msg: %s", xe.errCode, xe.errMsg)
}

// ErrCode get error code
func (xe *XError) ErrCode() uint32 {
	return xe.errCode
}

// ErrMsg get error msg
func (xe *XError) ErrMsg() string {
	return xe.errMsg
}

func NewCodeMsgError(code uint32, msg string) *XError {
	return &XError{
		errCode: code,
		errMsg:  msg,
	}
}

func NewCodeError(code uint32) *XError {
	return &XError{
		errCode: code,
		errMsg:  CodeMsg(code),
	}
}

func NewMsgError(msg string) *XError {
	return &XError{
		errCode: CodeServerCommError,
		errMsg:  msg,
	}
}
