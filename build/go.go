package build

import (
	"path"
	"runtime"

	"github.com/magefile/mage/sh"
)

var (
	DefaultGoEnv = map[string]string{"GO111MODULE": "on"}
)

type GoInstall struct {
	Env     map[string]string
	SrcPath string
}

type GoBuild struct {
	Env                      map[string]string
	SrcDir, OutDir, ExecName string
}

func (a *GoInstall) Run() error {
	return sh.RunWith(getGoEnv(a.Env), "go", "install", a.SrcPath)
}

func (a *GoBuild) Run() error {
	outDir := a.OutDir
	if outDir == "" {
		outDir = ".out"
	}
	execName := a.ExecName
	if execName == "" {
		execName = path.Base(a.SrcDir)
	}
	if runtime.GOOS == "windows" && path.Ext(execName) == "" {
		execName += ".exe"
	}
	return sh.RunWith(getGoEnv(a.Env), "go", "build", "-o", path.Join(a.OutDir, a.ExecName), a.SrcDir)
}

func getGoEnv(env map[string]string) map[string]string {
	if env != nil {
		return env
	}
	return DefaultGoEnv
}
