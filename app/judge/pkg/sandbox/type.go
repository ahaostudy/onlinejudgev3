package sandbox

import (
	"fmt"
	"github.com/ahaostudy/onlinejudge/app/judge/conf"
)

// Config sandbox config
type Config struct {
	*ExeConfig
	*FileConfig
	*LimitConfig
	*SafeConfig

	LogPath string // judger log path
}

type ExeConfig struct {
	ExePath string   // path of file to run
	Args    []string // arguments to run this process
	Env     []string // environment variables this process can get
}

type FileConfig struct {
	InputFile  string // redirect content of this file to process's stdin
	OutputFile string // redirect process's stdout to this file
	ErrorFile  string // redirect process's stderr to this file
}

type LimitConfig struct {
	MaxCPUTime           int   // max cpu time this process can cost, -1 for unlimited
	MaxRealTime          int   // max time this process can run, -1 for unlimited
	MaxMemory            int64 // max size of the process' virtual memory (address space), -1 for unlimited
	MaxStack             int64 // max size of the process' stack size
	MaxProcessNumber     int   // max number of processes that can be created for the real user id of the calling process, -1 for unlimited
	MaxOutputSize        int64 // max size of data this process can output to stdout, stderr and file, -1 for unlimited
	MemoryLimitCheckOnly int   // if this value equals 0, we will only check memory usage number, because setrlimit(maxrss) will cause some crash issues
}

type SafeConfig struct {
	SeccompRuleName string // seccomp rules used to limit process system calls
	UID             int    // user to run this process
	GID             int    // user group this process belongs to
}

func (c *Config) Arguments() []string {
	args := []string{
		fmt.Sprintf("--exe_path=%s", c.ExePath),
	}
	if c.LogPath != "" {
		args = append(args, fmt.Sprintf("--log_path=%s", c.LogPath))
	}
	if c.LimitConfig != nil {
		if c.MaxCPUTime != 0 {
			args = append(args, fmt.Sprintf("--max_cpu_time=%d", c.MaxCPUTime))
		}
		if c.MaxRealTime != 0 {
			args = append(args, fmt.Sprintf("--max_real_time=%d", c.MaxRealTime))
		}
		if c.MaxMemory != 0 {
			args = append(args, fmt.Sprintf("--max_memory=%d", c.MaxMemory))
		}
		if c.MaxStack != 0 {
			args = append(args, fmt.Sprintf("--max_stack=%d", c.MaxStack))
		}
		if c.MaxProcessNumber != 0 {
			args = append(args, fmt.Sprintf("--max_process_number=%d", c.MaxProcessNumber))
		}
		if c.MaxOutputSize != 0 {
			args = append(args, fmt.Sprintf("--max_output_size=%d", c.MaxOutputSize))
		}
		if c.MemoryLimitCheckOnly != 0 {
			args = append(args, fmt.Sprintf("--memory_limit_check_only=%d", c.MemoryLimitCheckOnly))
		}
	}
	if c.FileConfig != nil {
		if c.InputFile != "" {
			args = append(args, fmt.Sprintf("--input_file=%s", c.InputFile))
		}
		if c.OutputFile != "" {
			args = append(args, fmt.Sprintf("--output_file=%s", c.OutputFile))
		}
		if c.ErrorFile != "" {
			args = append(args, fmt.Sprintf("--error_file=%s", c.ErrorFile))
		}
	}
	if c.SafeConfig != nil {
		if c.SeccompRuleName != "" {
			args = append(args, fmt.Sprintf("--seccomp_rule_name=%s", c.SeccompRuleName))
		}
		if c.UID != 0 {
			args = append(args, fmt.Sprintf("--uid=%d", c.UID))
		}
		if c.GID != 0 {
			args = append(args, fmt.Sprintf("--gid=%d", c.GID))
		}
	}
	for _, arg := range c.Args {
		args = append(args, fmt.Sprintf("--args=%s", arg))
	}
	for _, env := range c.Env {
		args = append(args, fmt.Sprintf("--env=%s", env))
	}
	return args
}

func NewConfig(exeConf *ExeConfig, limitConf *LimitConfig, opts ...ConfigOption) *Config {
	cfg := &Config{
		ExeConfig:   exeConf,
		LogPath:     conf.GetConf().Sandbox.LogPath,
		LimitConfig: limitConf,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

type ConfigOption func(cfg *Config)

func WithFileConfig(c *FileConfig) ConfigOption {
	return func(cfg *Config) {
		cfg.FileConfig = c
	}
}

func WithSafeConfig(c *SafeConfig) ConfigOption {
	return func(cfg *Config) {
		cfg.SafeConfig = c
	}
}

func PatchLimitConfig(c *LimitConfig) {
	if c == nil {
		return
	}
	sandboxConf := conf.GetConf().Sandbox
	if c.MaxCPUTime == 0 {
		c.MaxCPUTime = sandboxConf.MaxCPUTime
	}
	if c.MaxRealTime == 0 {
		c.MaxRealTime = sandboxConf.MaxRealTime
	}
	if c.MaxMemory == 0 {
		c.MaxMemory = sandboxConf.MaxMemory
	}
	if c.MaxStack == 0 {
		c.MaxStack = sandboxConf.MaxStack
	}
	if c.MaxProcessNumber == 0 {
		c.MaxProcessNumber = sandboxConf.MaxProcessNumber
	}
	if c.MaxOutputSize == 0 {
		c.MaxOutputSize = sandboxConf.MaxOutputSize
	}
	if c.MemoryLimitCheckOnly == 0 {
		c.MemoryLimitCheckOnly = sandboxConf.MemoryLimitCheckOnly
	}
}

// Result sandbox response result
type Result struct {
	CpuTime  int64 `json:"cpu_time"`
	RealTime int64 `json:"real_time"`
	Memory   int64 `json:"memory"`
	Signal   int   `json:"signal"`
	ExitCode int   `json:"exit_code"`
	Error    int   `json:"error"`
	Result   int   `json:"result"`
}

const (
	ResultSuccess = iota
	ResultCpuTimeLimitExceeded
	ResultRealTimeLimitExceeded
	ResultMemoryLimitExceeded
	ResultRuntimeError
	ResultSystemError
)
