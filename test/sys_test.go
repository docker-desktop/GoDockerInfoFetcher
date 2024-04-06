package test

import (
	"testing"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

func TestIsDockerInstalled(t *testing.T) {
	installed, err := pkg.IsDockerInstalled()
	if err != nil {
		t.Errorf("IsDockerInstalled() error = %v, wantErr false", err)
		return
	}

	if !installed {
		t.Errorf("IsDockerInstalled() = %v, want true", installed)
	}
}

func TestGetDockerVersion(t *testing.T) {
	version := pkg.GetDockerVersion()
	if version == "" {
		t.Errorf("GetDockerVersion() = %v, want not empty", version)
	}
}
