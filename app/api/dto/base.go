package dto

import (
	"fmt"

	"github.com/cloudwego/kitex/pkg/kerrors"
)

type BaseResp struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

const (
	SuccessCode       = 0
	ServerFailureCode = 50000
)

const (
	SuccessMsg       = "ok"
	ServerFailureMsg = "server failure: %s"
)

func SuccessResp() BaseResp {
	return BaseResp{StatusCode: SuccessCode, StatusMsg: SuccessMsg}
}

func BuildBaseResp(err error) BaseResp {
	if err == nil {
		return SuccessResp()
	}

	if bizErr, ok := kerrors.FromBizStatusError(err); ok {
		return BaseResp{StatusCode: bizErr.BizStatusCode(), StatusMsg: bizErr.BizMessage()}
	}

	return BaseResp{StatusCode: ServerFailureCode, StatusMsg: fmt.Sprintf(ServerFailureMsg, err.Error())}
}
