// +build mage

package main

import (
	"errors"
	"github.com/andrey-ko/simple-rest/build"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Aliases = map[string]interface{}{
	"build-image":  BuildImage,
	"build-image-centos":  BuildImageCentos,
	"build-image-win1809": BuildImageWin1809,
	"build-image-win1803": BuildImageWin1803,
	"build-image-win1709": BuildImageWin1709,
}

func Build() error {
	action := &build.GoBuild{
		SrcDir: "./cmd/simple-rest",
	}
	return action.Run()
}

func buildImage(target string) error {
	action := &build.DockerBuild{
		DockerFile: "Dockerfile." + target,
		Tag:        "akolomentsev/simple-rest:" + target,
	}
	return action.Run()
}
func BuildImage() error {
	args,err := build.ParseStdin()
	if err != nil {
		return err
	}
	target, ok := args["target"]
	if !ok || target=="" {
		return errors.New("target is missing")
	}
	return buildImage(target)
}
func BuildImageWin1809() error {
	return buildImage("win1809")
}
func BuildImageWin1803() error {
	return buildImage("win1803")
}
func BuildImageWin1709() error {
	return buildImage("win1709")
}
func BuildImageCentos() error {
	return buildImage("centos")
}

func InstallGoTools() error {
	tools := []string{
		"golang.org/x/lint/golint",
		"github.com/gordonklaus/ineffassign",
		"github.com/client9/misspell/cmd/misspell",
		"honnef.co/go/tools/cmd/staticcheck",
	}
	for _, t := range tools {
		action := &build.GoInstall{SrcPath: t}
		if err := action.Run(); err != nil {
			return err
		}
	}
	return nil
}

func Check() error {
	mg.Deps(Build)
	mg.Deps(InstallGoTools)

	for _, dir := range []string{"./cmd/", "./build/"} {
		if err := sh.Run("gofmt", "-s", "-w", dir); err != nil {
			return err
		}
		if err := sh.Run("goimports", "-w", dir); err != nil {
			return err
		}
		if err := sh.Run("staticcheck", "-checks", "all",  dir +"..."); err != nil {
			return err
		}
	}

	return nil
}


