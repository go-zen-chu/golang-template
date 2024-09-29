//go:build mage
// +build mage

package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/go-zen-chu/aictl/internal/mage"
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
	outMsg, errMsg, err := mage.RunLongRunningCmdWithLog("go install go.uber.org/mock/mockgen@latest")
	if err != nil {
		return fmt.Errorf("installing mockgen: %w\nstdout: %s\nstderr: %s\n", err, outMsg, errMsg)
	}
	return nil
}

/*=======================
workflow
=======================*/

func DockerLogin() error {
	return mage.DockerLogin()
}

func DockerBuildLatest() error {
	return mage.DockerBuildLatest(imageRegistry, repository, dockerFileLocation)
}

func DockerPublishLatest() error {
	return mage.DockerPublishLatest(imageRegistry, repository)
}

func DockerBuildPublishWithGenTag() error {
	return mage.DockerBuildPublishGeneratedImageTag(imageRegistry, repository, dockerFileLocation)
}
