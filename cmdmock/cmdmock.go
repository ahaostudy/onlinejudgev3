package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// StringSlice is a custom flag type to capture command line arguments into a slice of strings
type StringSlice []string

// String is needed to satisfy interface requirement
func (s *StringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set appends new value to the flag
func (s *StringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func main() {
	var args, env StringSlice
	exePath := flag.String("exe_path", "", "")
	inputFile := flag.String("input_file", "", "")
	outputFile := flag.String("output_file", "", "")
	errorFile := flag.String("error_file", "", "")
	flag.Var(&args, "args", "")
	// other
	flag.String("log_path", "", "")
	flag.Int64("max_cpu_time", 0, "")
	flag.Int64("max_real_time", 0, "")
	flag.Int64("max_memory", 0, "")
	flag.Int64("max_stack", 0, "")
	flag.Int64("max_process_number", 0, "")
	flag.Int64("max_output_size", 0, "")
	flag.Int64("memory_limit_check_only", 0, "")
	flag.String("seccomp_rule_name", "", "")
	flag.Int64("uid", 0, "")
	flag.Int64("gid", 0, "")
	flag.Var(&env, "env", "")
	flag.Parse()

	result := &Result{}

	if *exePath == "" {
		result.Result = 5
		result.Stdout("exe_path is required")
	}

	cmd := exec.Command(*exePath, args...)
	if *inputFile != "" {
		input, err := os.Open(*inputFile)
		if err != nil {
			result.Result = 5
			result.Stdout("open file error: " + err.Error())
		}
		defer input.Close()
		cmd.Stdin = input
	}

	var stdout, stderr bytes.Buffer

	if *outputFile != "" {
		output, err := os.Create(*outputFile)
		if err != nil {
			result.Result = 5
			result.Stdout("create file error: " + err.Error())
		}
		defer output.Close()
		cmd.Stdout = io.MultiWriter(&stdout, output)
	}

	if *errorFile != "" {
		errorOut, err := os.Create(*errorFile)
		if err != nil {
			result.Result = 5
			result.Stdout("create file error: " + err.Error())
		}
		defer errorOut.Close()
		cmd.Stderr = io.MultiWriter(&stderr, errorOut)
	}

	startTime := time.Now()
	var memStart runtime.MemStats
	runtime.ReadMemStats(&memStart)

	err := cmd.Run()

	endTime := time.Now()
	var memEnd runtime.MemStats
	runtime.ReadMemStats(&memEnd)

	execTime := endTime.Sub(startTime)
	memUsed := memEnd.Alloc - memStart.Alloc
	result.CpuTime = int64(execTime / time.Millisecond)
	result.RealTime = int64(execTime / time.Millisecond)
	result.Memory = int64(memUsed)

	if err != nil {
		result.Result = 4
		result.Stdout(fmt.Sprintf("cmd run failed: %s, stdout: %s, stderr: %s\n", err.Error(), stdout.String(), stderr.String()))
	}

	result.Stdout("")
}

type Result struct {
	CpuTime  int64  `json:"cpu_time"`
	RealTime int64  `json:"real_time"`
	Memory   int64  `json:"memory"`
	Signal   int    `json:"signal"`
	ExitCode int    `json:"exit_code"`
	Error    int    `json:"error"`
	Result   int    `json:"result"`
	Message  string `json:"message"`
}

func (r *Result) Str() string {
	bs, err := json.Marshal(r)
	if err != nil {
		return `{"cpu_time":0,"signal":0,"memory":0,"exit_code":0,"result":5,"error":0,"real_time":0,"message":"json marshal error"}`
	}
	return string(bs)
}

func (r *Result) Stdout(msg string) {
	r.Message = msg
	_, _ = fmt.Fprintln(os.Stdout, r.Str())
	os.Exit(0)
}
