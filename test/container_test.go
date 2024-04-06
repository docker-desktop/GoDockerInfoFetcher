package test

import (
	"testing"

	con "github.com/docker-desktop/GoDockerInfoFetcher/pkg/container"
)

func TestGetDockerContainers(t *testing.T) {
	out, err := con.GetDockerContainers()
	if err != nil {
		t.Errorf("GetDockerContainers() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetDockerContainers() returned empty list, expected at least one container")
	}
}

func TestGetRunningDockerContainers(t *testing.T) {
	out, err := con.GetRunningDockerContainers()
	if err != nil {
		t.Errorf("GetRunningDockerContainers() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetRunningDockerContainers() returned empty list, expected at least one running container")
	}
}
