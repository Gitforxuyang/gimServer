package err

import (
	"encoding/json"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GimError struct {
	Code   int32
	Msg    string
	Detail string //错误详情
}

func (m *GimError) Error() string {
	return fmt.Sprintf("err code:%d msg:%s detail:%s", m.Code, m.Msg, m.Detail)
}

func (m *GimError) SetDetail(detail string) *GimError {
	return &GimError{Code: m.Code, Msg: m.Msg, Detail: detail}
}

var (
	UserNotFoundError = &GimError{Code: 3001, Msg: "未找到用户"}
	ParamsError       = &GimError{Code: 3002, Msg: "参数错误"}
	UnknownError      = &GimError{Code: 3000, Msg: "未知错误"}
)

func EncodeStatus(e *GimError) *status.Status {
	status := status.New(codes.Code(e.Code), e.Error())

	return status
}

func DecodeStatus(e error) *GimError {

	status, ok := status.FromError(e)

	if !ok {
		return Parse(e.Error())
	} else {
		return Parse(status.Message())
	}
}

func Parse(err string) *GimError {
	e := new(GimError)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		a := UnknownError.SetDetail(err)
		return a
	}
	return e
}

func FromError(err error) *GimError {
	if verr, ok := err.(*GimError); ok && verr != nil {
		return verr
	}

	return Parse(err.Error())
}
