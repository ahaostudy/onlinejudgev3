package osfile

import (
	"os"
	"path/filepath"

	"github.com/ahaostudy/onlinejudge/app/judge/conf"
	"github.com/google/uuid"
)

const tmpPath = "tmp"

// AbsPath use relative path to splice the base path
// except for this method, the parameters and return value of other methods are all relative paths
func AbsPath(path ...string) string {
	return filepath.Join(conf.GetConf().File.BaseUrl, filepath.Join(path...))
}

func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(AbsPath(path))
}

func WriteFile(path string, raw []byte) (err error) {
	absPath := AbsPath(path)
	err = os.MkdirAll(filepath.Dir(absPath), 0o777)
	if err != nil {
		return err
	}
	err = os.WriteFile(absPath, raw, 0o666)
	return
}

func RemoveAll(path string) error {
	return os.RemoveAll(AbsPath(path))
}

func MakeDir(dir string) (*Dir, error) {
	path := filepath.Join(dir, uuid.New().String())
	err := os.MkdirAll(AbsPath(path), 0o777)
	if err != nil {
		return nil, err
	}
	return &Dir{BaseUrl: path}, nil
}

func MakeTmpDir() (*Dir, error) {
	return MakeDir(tmpPath)
}

type Dir struct {
	BaseUrl string
}

func (dir *Dir) Path(path ...string) string {
	return filepath.Join(dir.BaseUrl, filepath.Join(path...))
}

func (dir *Dir) WriteFile(filename string, raw []byte) (string, error) {
	path := dir.Path(filename)
	err := WriteFile(path, raw)
	if err != nil {
		return "", err
	}
	return path, nil
}

func (dir *Dir) CleanUp() {
	_ = RemoveAll(dir.Path())
}
