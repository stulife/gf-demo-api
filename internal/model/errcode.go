package model

import (
	"fmt"
)

type ErrCode int

const (
	CodeOK             ErrCode = 200
	UnAuthorized       ErrCode = 401
	CodeNotFound       ErrCode = 404
	CodeInternalError  ErrCode = 500
	CodeUnknown        ErrCode = 520
	CodeBusinessFailed ErrCode = 10000
)

var errCodeName = map[ErrCode]string{
	CodeOK:             "Success",
	UnAuthorized:       "Not Authorized",
	CodeInternalError:  "INTERNAL ERROR",
	CodeNotFound:       "Resource Does Not Exist",
	CodeUnknown:        "Unknown Error",
	CodeBusinessFailed: "业务出错",
}

func (e ErrCode) Desc() string {
	if s, ok := errCodeName[e]; ok {
		return s
	}
	return fmt.Sprintf("unknown error code 0x%x", uint32(e))
}

func (e ErrCode) Code() int {
	return int(e)
}
