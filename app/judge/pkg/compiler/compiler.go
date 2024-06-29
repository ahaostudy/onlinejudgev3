package compiler

import (
	"bytes"
	"fmt"
	"github.com/ahaostudy/onlinejudge/app/judge/conf"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/language"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/osfile"
	"github.com/ahaostudy/onlinejudge/app/judge/pkg/sandbox"
	"github.com/cloudwego/kitex/pkg/klog"
	"os/exec"
	"path/filepath"
	"strings"
)

type Compiler func(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error)

type CompileError struct {
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
	Err    error
}

func (e *CompileError) Error() string {
	return fmt.Sprintf("compile error, error: %s, stdout: %s, stderr: %s", e.Err.Error(), e.Stdout.String(), e.Stderr.String())
}

func NewCompileError(stdout, stderr *bytes.Buffer, err error) *CompileError {
	return &CompileError{Stdout: stdout, Stderr: stderr, Err: err}
}

func GetCompiler(lang language.Language) (Compiler, error) {
	switch lang {
	case language.Language_CPP:
		return CompileCPP, nil
	case language.Language_C:
		return CompileC, nil
	case language.Language_Python3:
		return CompilePython3, nil
	case language.Language_Java:
		return CompileJava, nil
	case language.Language_Go:
		return CompileGo, nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", lang)
	}
}

func CompileCPP(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error) {
	exePath := osfile.AbsPath(dir.Path(ExeName(codePath)))
	stdout, stderr, err := Exec(conf.GetConf().Compiler.CPP, "-std=c++11", osfile.AbsPath(codePath), "-o", exePath)
	if err != nil {
		return nil, NewCompileError(stdout, stderr, err)
	}
	return &sandbox.ExeConfig{ExePath: exePath}, nil
}

func CompileC(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error) {
	exePath := osfile.AbsPath(dir.Path(ExeName(codePath)))
	stdout, stderr, err := Exec(conf.GetConf().Compiler.C, "-std=c11", osfile.AbsPath(codePath), "-o", exePath)
	if err != nil {
		return nil, NewCompileError(stdout, stderr, err)
	}
	return &sandbox.ExeConfig{ExePath: exePath}, nil
}

func CompilePython3(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error) {
	return &sandbox.ExeConfig{
		ExePath: conf.GetConf().Compiler.Python3,
		Args:    []string{osfile.AbsPath(codePath)},
	}, nil
}

func CompileJava(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error) {
	exeDir, exeName := osfile.AbsPath(dir.Path()), ExeName(codePath)
	stdout, stderr, err := Exec(conf.GetConf().Compiler.JavaC, "-d", exeDir, "-encoding", "UTF8", osfile.AbsPath(codePath))
	if err != nil {
		return nil, NewCompileError(stdout, stderr, err)
	}
	exeConf := &sandbox.ExeConfig{
		ExePath: conf.GetConf().Compiler.Java,
		Args:    []string{"-cp", exeDir, fmt.Sprintf("-XX:MaxRAM=%d", limitConf.MaxMemory*4), "-Djava.security.manager", "-Dfile.encoding=UTF-8", "-Djava.awt.headless=true", exeName},
	}
	limitConf.MaxMemory = 0
	limitConf.MemoryLimitCheckOnly = 1
	return exeConf, nil
}

func CompileGo(codePath string, limitConf *sandbox.LimitConfig, dir *osfile.Dir) (*sandbox.ExeConfig, error) {
	exePath := osfile.AbsPath(dir.Path(ExeName(codePath)))
	stdout, stderr, err := Exec(conf.GetConf().Compiler.Go, "build", "-o", exePath, osfile.AbsPath(codePath))
	if err != nil {
		return nil, NewCompileError(stdout, stderr, err)
	}
	return &sandbox.ExeConfig{ExePath: exePath}, nil
}

func ExeName(codePath string) string {
	return strings.TrimSuffix(filepath.Base(codePath), filepath.Ext(codePath))
}

func Exec(name string, arg ...string) (*bytes.Buffer, *bytes.Buffer, error) {
	cmd := exec.Command(name, arg...)
	klog.Infof("[compile] %v\n", cmd.String())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	klog.Debugf("[compile] stdout: %v\n", stdout.String())
	klog.Debugf("[compile] stderr: %v\n", stderr.String())
	return &stdout, &stderr, err
}
