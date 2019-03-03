package test

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

var cachedOrbsKeyGeneratorBinaryPath string

type orbsKeyGenerator struct {
}

func compileOrbsKeyGenerator() string {
	if orbsKeyGeneratorPathFromEnv := os.Getenv("ORBS_KEY_GENERATOR_PATH"); orbsKeyGeneratorPathFromEnv != "" {
		return orbsKeyGeneratorPathFromEnv
	}

	if cachedOrbsKeyGeneratorBinaryPath != "" {
		return cachedOrbsKeyGeneratorBinaryPath // cache compilation once per process
	}

	tempDir, err := ioutil.TempDir("", "orbskg")
	if err != nil {
		panic(err)
	}

	binaryOutputPath := tempDir + "/orbs-key-generator"
	goCmd := path.Join(runtime.GOROOT(), "bin", "go")
	cmd := exec.Command(goCmd, "build", "-o", binaryOutputPath, ".")
	cmd.Dir = path.Join(getCurrentSourceFileDirPath(), "..")
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("compilation failed: %s\noutput:\n%s\n", err.Error(), out))
	} else {
		fmt.Printf("compiled orbs-key-generator successfully:\n %s\n", binaryOutputPath)
	}

	cachedOrbsKeyGeneratorBinaryPath = binaryOutputPath
	return cachedOrbsKeyGeneratorBinaryPath
}

func (g *orbsKeyGenerator) Run(args ...string) (string, error) {
	out, err := exec.Command(compileOrbsKeyGenerator(), args...).CombinedOutput()
	return string(out), err
}

func OrbsKeyGenerator() *orbsKeyGenerator {
	return &orbsKeyGenerator{}
}

func getCurrentSourceFileDirPath() string {
	_, filename, _, _ := runtime.Caller(1)
	return filepath.Dir(filename)
}
