package main

import (
	"context"
	"github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc"
	judgeservice "github.com/ahaostudy/onlinejudge/kitex_gen/judgesvc/judgeservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"os"
)

var (
	code = []byte(`
	#include <stdio.h>
	int a, b;
	int main() {
		scanf("%d%d", &a, &b);
		printf("%d\n", a + b);
		return 0;
	}
	`)
	lang   = judgesvc.Language_C
	input  = []byte("3432 12")
	fileId = "7cbb450a-c859-4c92-95b2-d4ccaa36f340"
)

func main() {
	cli := judgeservice.MustNewClient("judge",
		client.WithHostPorts("127.0.0.1:8881"),
		client.WithTransportProtocol(transport.TTHeaderFramed),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	ctx := context.Background()

	uploadResp, err := cli.UploadCode(ctx, &judgesvc.UploadCodeRequest{
		Code:     code,
		Language: lang,
	})
	handleError(err)
	klog.Info("FileId: ", uploadResp.FileId)

	judgeResp, err := cli.Judge(ctx, &judgesvc.JudgeRequest{
		//Code: code,
		CodeFileId: &uploadResp.FileId,
		Input:      input,
		Language:   lang,
	})
	handleError(err)
	klog.Infof("Status: %v, Result: %#v\n", judgeResp.Status.String(), judgeResp)

	_, err = cli.DeleteCode(ctx, &judgesvc.DeleteCodeRequest{FileId: fileId})
	handleError(err)
	_, err = cli.DeleteCode(ctx, &judgesvc.DeleteCodeRequest{FileId: uploadResp.FileId})
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		klog.Error(err)
		os.Exit(1)
	}
}
