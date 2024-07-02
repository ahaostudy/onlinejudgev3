package main

import (
	"context"
	"github.com/ahaostudy/onlinejudge/kitex_gen/base"

	"github.com/ahaostudy/onlinejudge/app/judge/handler"
	judgesvc "github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc"
)

// JudgeServiceImpl implements the last service interface defined in the IDL.
type JudgeServiceImpl struct{}

// Judge implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) Judge(ctx context.Context, req *judgesvc.JudgeRequest) (resp *judgesvc.JudgeResponse, err error) {
	return handler.Judge(ctx, req)
}

// UploadCode implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) UploadCode(ctx context.Context, req *judgesvc.UploadCodeRequest) (resp *judgesvc.UploadCodeResponse, err error) {
	return handler.UploadCode(ctx, req)
}

// DeleteCode implements the JudgeServiceImpl interface.
func (s *JudgeServiceImpl) DeleteCode(ctx context.Context, req *judgesvc.DeleteCodeRequest) (resp *base.Empty, err error) {
	return handler.DeleteCode(ctx, req)
}
