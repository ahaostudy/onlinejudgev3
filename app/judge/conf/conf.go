package conf

import (
	"path/filepath"

	ktconf "github.com/ahaostudy/kitextool/conf"
	"github.com/kr/pretty"
)

var conf = new(Conf)

func GetConf() *Conf {
	return conf
}

func init() {
	ktconf.LoadFiles(conf,
		filepath.Join("conf", "conf.yaml"),
		filepath.Join("app", "judge", "conf", "conf.yaml"),
	)
	_, _ = pretty.Printf("%+v\n", conf)
}

type Conf struct {
	ktconf.ServerConf
	Sandbox  Sandbox  `yaml:"sandbox"`
	Compiler Compiler `yaml:"compiler"`
	File     File     `yaml:"file"`
}

type Sandbox struct {
	ExePath              string `yaml:"exe_path"`
	LogPath              string `yaml:"log_path"`
	MaxCPUTime           int    `yaml:"max_cpu_time"`
	MaxRealTime          int    `yaml:"max_real_time"`
	MaxMemory            int64  `yaml:"max_memory"`
	MaxStack             int64  `yaml:"max_stack"`
	MaxProcessNumber     int    `yaml:"max_process_number"`
	MaxOutputSize        int64  `yaml:"max_output_size"`
	MemoryLimitCheckOnly int    `yaml:"memory_limit_check_only"`
}

type Compiler struct {
	C       string `yaml:"c"`
	CPP     string `yaml:"cpp"`
	Go      string `yaml:"go"`
	Java    string `yaml:"java"`
	JavaC   string `yaml:"javac"`
	Python3 string `yaml:"python3"`
}

type File struct {
	BaseUrl string `yaml:"base_url"`
}
