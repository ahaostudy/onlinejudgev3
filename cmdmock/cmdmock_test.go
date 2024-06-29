package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	src = "./cmdmock.go"

	testDir    = "./test"
	codeFile   = filepath.Join(testDir, "main.go")
	inputFile  = filepath.Join(testDir, "input.in")
	outputFile = filepath.Join(testDir, "output.out")
	errorFile  = filepath.Join(testDir, "error.err")
	exePath    = strings.TrimSuffix(codeFile, filepath.Ext(codeFile))
)

const (
	code = `package main

import "fmt"

func main() {
	var a, b int
	_, err := fmt.Scan(&a, &b)
	if err != nil {
		panic(err)
	}
	fmt.Println(a + b)
}
`
	input  = "1 2"
	output = "3"
)

func PrepareTestData() error {
	err := os.MkdirAll(testDir, 0o777)
	if err != nil {
		return err
	}
	err = os.WriteFile(inputFile, []byte(input), 0o666)
	if err != nil {
		return err
	}
	err = os.WriteFile(codeFile, []byte(code), 0o666)
	if err != nil {
		return err
	}
	err = exec.Command("go", "build", "-o", exePath, codeFile).Run()
	if err != nil {
		return err
	}
	return nil
}

func CheckOutput() (bool, error) {
	content, err := os.ReadFile(outputFile)
	if err != nil {
		return false, err
	}
	return strings.TrimSpace(string(content)) == output, nil
}

func CleanUpTestData() {
	_ = os.RemoveAll(testDir)
}

func TestCmd(t *testing.T) {
	err := PrepareTestData()
	assert.NoError(t, err)
	defer CleanUpTestData()

	cmd := exec.Command("go", "run", src,
		f("--exe_path=%s", exePath),
		f("--input_file=%s", inputFile),
		f("--output_file=%s", outputFile),
		f("--error_file=%s", errorFile),
	)
	var stdout, stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	err = cmd.Run()
	assert.NoError(t, err)
	assert.Equal(t, stderr.Len(), 0)

	result := new(Result)
	err = json.Unmarshal(stdout.Bytes(), result)
	assert.NoError(t, err)
	t.Logf("%+v\n", result)

	ok, err := CheckOutput()
	assert.NoError(t, err)
	assert.True(t, ok)
}

func f(f string, a ...any) string {
	return fmt.Sprintf(f, a...)
}
