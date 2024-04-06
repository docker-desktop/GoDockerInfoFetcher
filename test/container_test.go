package test

import (
	"testing"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

var (
	dockerContainerFetcher = pkg.NewDockerContainerFetcher()
)

func TestGetAllContainers(t *testing.T) {
	out, err := dockerContainerFetcher.GetAllContainers()
	if err != nil {
		t.Errorf("GetDockerContainers() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetDockerContainers() returned empty list, expected at least one container")
	}
}

func TestGetRunningContainers(t *testing.T) {
	out, err := dockerContainerFetcher.GetRunningContainers()
	if err != nil {
		t.Errorf("GetRunningDockerContainers() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetRunningDockerContainers() returned empty list, expected at least one running container")
	}
}
