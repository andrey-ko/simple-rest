package build

import (
	"github.com/magefile/mage/sh"
)

var (
	DefaultDockerEnv  = map[string]string{}
	DefaultDockerOpts = &DockerOpts{}
)

type DockerOpts struct {
	Host string
}

type DockerBuild struct {
	Env map[string]string
	*DockerOpts
	Tag        string
	DockerFile string
	ContextDir string
}

type DockerPush struct {
	Env map[string]string
	*DockerOpts
	Image string
}

func (a *DockerBuild) Run() error {
	args := getDockerOpts(a.DockerOpts, "build")
	if a.Tag != "" {
		args = append(args, "-t", a.Tag)
	}
	if a.DockerFile != "" {
		args = append(args, "-f", a.DockerFile)
	}
	if a.ContextDir != "" {
		args = append(args, a.ContextDir)
	} else {
		args = append(args, ".")
	}
	return sh.RunWith(getDockerEnv(a.Env), "docker", args...)
}

func (a *DockerPush) Run() error {
	args := getDockerOpts(a.DockerOpts, "push", a.Image)
	return sh.RunWith(getDockerEnv(a.Env), "docker", args...)
}

func getDockerEnv(env map[string]string) map[string]string {
	if env != nil {
		return env
	}
	return DefaultDockerEnv
}

func getDockerOpts(opts *DockerOpts, args ...string) []string {
	var res []string
	if opts == nil {
		opts = DefaultDockerOpts
	}
	if opts.Host != "" {
		res = append(res, "-H", opts.Host)
	}
	return append(res, args...)
}
