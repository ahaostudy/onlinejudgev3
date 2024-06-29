package sandbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/ahaostudy/onlinejudge/app/judge/conf"
	"github.com/cloudwego/kitex/pkg/klog"
)

type SandBox struct {
	Config *Config
}

func NewSandbox(config *Config) *SandBox {
	return &SandBox{Config: config}
}

func (s *SandBox) Exec() (*Result, error) {
	args := s.Config.Arguments()
	cmd := exec.Command(conf.GetConf().Sandbox.ExePath, args...)
	klog.Infof("[sandbox] %v\n", cmd.String())

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	klog.Debugf("[sandbox] stdout: %v\n", stdout.String())
	klog.Debugf("[sandbox] stderr: %v\n", stderr.String())
	if err != nil {
		return nil, fmt.Errorf("failed to execute sandbox command: %w", err)
	}

	// parse the execution result
	result := new(Result)
	err = json.Unmarshal(stdout.Bytes(), result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal sandbox result: %w", err)
	}

	return result, nil
}
