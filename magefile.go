//go:build mage
// +build mage

package main

import (
	"fmt"
	"log/slog"
	"os"

	gbt "github.com/go-zen-chu/go-build-tools"
)

const imageRegistry = "amasuda"
const repository = "aictl"
const dockerFileLocation = "."

func init() {
	// by default, magefile does not output stderr
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

/*=======================
setup
=======================*/

// InstallDevTools installs required development tools for this project
func InstallDevTools() error {
	outMsg, errMsg, err := gbt.RunLongRunningCmdWithLog("go install go.uber.org/mock/mockgen@latest")
	if err != nil {
		return fmt.Errorf("installing mockgen: %w\nstdout: %s\nstderr: %s\n", err, outMsg, errMsg)
	}
	outMsg, errMsg, err = gbt.RunLongRunningCmdWithLog("brew install golangci-lint")
	if err != nil {
		return fmt.Errorf("installing mockgen: %w\nstdout: %s\nstderr: %s\n", err, outMsg, errMsg)
	}
	return nil
}

/*=======================
workflow
=======================*/

func DockerLogin() error {
	return gbt.DockerLogin()
}

func DockerBuildLatest() error {
	return gbt.DockerBuildLatest(imageRegistry, repository, dockerFileLocation)
}

func DockerPublishLatest() error {
	return gbt.DockerPublishLatest(imageRegistry, repository)
}

func DockerBuildPublishWithGenTag() error {
	return gbt.DockerBuildPublishGeneratedImageTag(imageRegistry, repository, dockerFileLocation)
}
