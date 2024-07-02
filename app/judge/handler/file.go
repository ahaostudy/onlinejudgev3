package handler

import (
	"context"
	"github.com/ahaostudy/onlinejudge/kitex_gen/base"
	"path/filepath"

	"github.com/ahaostudy/onlinejudge/app/judge/dal/cache"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/language"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/osfile"
	"github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

func UploadCode(ctx context.Context, req *judgesvc.UploadCodeRequest) (resp *judgesvc.UploadCodeResponse, err error) {
	lang, err := language.FromString(req.Language.String())
	if err != nil {
		return nil, kerrors.NewBizStatusError(40021, "language is not exist")
	}
	dir, err := osfile.MakeTmpDir()
	if err != nil {
		return nil, kerrors.NewBizStatusError(50021, "failed to create tmp dir: "+err.Error())
	}
	codePath, err := dir.WriteFile(lang.FileName(), req.Code)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50022, "failed to write file: "+err.Error())
	}
	fileId, err := cache.NewFileCache(ctx).Set(codePath)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50023, "failed to cache store: "+err.Error())
	}
	resp = &judgesvc.UploadCodeResponse{FileId: fileId}
	return
}

func DeleteCode(ctx context.Context, req *judgesvc.DeleteCodeRequest) (resp *base.Empty, err error) {
	path, err := cache.NewFileCache(ctx).Get(req.FileId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(40031, "file is not exist")
	}
	err = osfile.RemoveAll(filepath.Dir(path))
	if err != nil {
		return nil, kerrors.NewBizStatusError(40031, "file is not exist")
	}
	err = cache.NewFileCache(ctx).Del(req.FileId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50031, "failed to delete cache")
	}
	return
}
