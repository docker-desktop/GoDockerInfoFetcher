package test

import (
	"testing"

	GoDockerInfoFetcher "github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

func TestIsDockerInstalled(t *testing.T) {
	installed, err := GoDockerInfoFetcher.IsDockerInstalled()
	if err != nil {
		t.Errorf("IsDockerInstalled() error = %v, wantErr false", err)
		return
	}

	if !installed {
		t.Errorf("IsDockerInstalled() = %v, want true", installed)
	}
}

func TestGetDockerVersion(t *testing.T) {
	version := GoDockerInfoFetcher.GetDockerVersion()
	if version == "" {
		t.Errorf("GetDockerVersion() = %v, want not empty", version)
	}
}
