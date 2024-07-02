package handler

import (
	"context"
	"errors"

	"github.com/ahaostudy/onlinejudge/app/judge/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/compiler"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/language"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/osfile"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/sandbox"
	"github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/klog"
)

const (
	InputFileName  = "in.in"
	OutputFileName = "out.out"
	ErrorFileName  = "err.err"
)

func Judge(ctx context.Context, req *judgesvc.JudgeRequest) (resp *judgesvc.JudgeResponse, err error) {
	lang, err := language.FromString(req.Language.String())
	if err != nil {
		return nil, kerrors.NewBizStatusError(40011, "language is not exist")
	}

	namespace, err := osfile.MakeTmpDir()
	if err != nil {
		return nil, kerrors.NewBizStatusError(50011, "failed to make directory: "+err.Error())
	}
	defer namespace.CleanUp()

	var codePath string
	if req.CodeFileId != nil {
		codePath, err = cache.NewFileCache(ctx).Get(*req.CodeFileId)
		if err != nil {
			return nil, kerrors.NewBizStatusError(40012, "invalid code_path_id")
		}
	} else {
		if req.Code == nil {
			return nil, kerrors.NewBizStatusError(40013, "code is empty")
		}
		codePath, err = namespace.WriteFile(lang.FileName(), req.Code)
		if err != nil {
			return nil, kerrors.NewBizStatusError(50012, "failed to create code file: "+err.Error())
		}
	}

	// set sandbox execution parameters
	limitConfig := &sandbox.LimitConfig{}
	if req.Limit != nil {
		limit := req.Limit
		limitConfig.MaxCPUTime = int(limit.MaxCPUTime)
		limitConfig.MaxRealTime = int(limit.MaxRealTime)
		limitConfig.MaxMemory = limit.MaxMemory
		limitConfig.MaxStack = limit.MaxStack
		limitConfig.MaxProcessNumber = int(limit.MaxProcessNumber)
		limitConfig.MaxOutputSize = limit.MaxOutputSize
		limitConfig.MemoryLimitCheckOnly = int(limit.MemoryLimitCheckOnly)
	}
	sandbox.PatchLimitConfig(limitConfig)

	// compiler code
	compile, err := compiler.GetCompiler(lang)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40013, err.Error())
	}
	exeConf, err := compile(codePath, limitConfig, namespace)
	if err != nil {
		klog.Errorf("[judge] compiler error: %v", err.Error())
		var ce *compiler.CompileError
		if errors.As(err, &ce) {
			return &judgesvc.JudgeResponse{Status: judgesvc.JudgeStatus_CompileError, Output: ce.Stdout.String(), Error: ce.Stderr.String()}, nil
		}
		return &judgesvc.JudgeResponse{Status: judgesvc.JudgeStatus_CompileError}, nil
	}

	// prepare input and output files
	inputPath, err := namespace.WriteFile(InputFileName, req.Input)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50013, "failed to create input file: "+err.Error())
	}
	outputPath := namespace.Path(OutputFileName)
	errorPath := namespace.Path(ErrorFileName)
	box := sandbox.NewSandbox(sandbox.NewConfig(
		exeConf,
		limitConfig,
		sandbox.WithFileConfig(&sandbox.FileConfig{
			InputFile:  osfile.AbsPath(inputPath),
			OutputFile: osfile.AbsPath(outputPath),
			ErrorFile:  osfile.AbsPath(errorPath),
		}),
	))

	result, err := box.Exec()
	if err != nil {
		return nil, kerrors.NewBizStatusError(50014, "failed to sandbox execute: "+err.Error())
	}
	var status judgesvc.JudgeStatus
	switch result.Result {
	case sandbox.ResultSuccess:
		status = judgesvc.JudgeStatus_Finished
	case sandbox.ResultCpuTimeLimitExceeded, sandbox.ResultRealTimeLimitExceeded:
		status = judgesvc.JudgeStatus_TimeLimitExceeded
	case sandbox.ResultMemoryLimitExceeded:
		status = judgesvc.JudgeStatus_MemoryLimitExceeded
	case sandbox.ResultRuntimeError:
		status = judgesvc.JudgeStatus_RuntimeError
	case sandbox.ResultSystemError:
		return nil, kerrors.NewBizStatusError(50015, "program execution failure")
	}

	output, err := osfile.ReadFile(outputPath)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50016, "failed to read output: "+err.Error())
	}
	errorOutput, err := osfile.ReadFile(errorPath)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50017, "failed to read error output: "+err.Error())
	}

	resp = &judgesvc.JudgeResponse{
		Time:   result.CpuTime,
		Memory: result.Memory,
		Status: status,
		Output: string(output),
		Error:  string(errorOutput),
	}
	return
}
