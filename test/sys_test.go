package test

import (
	"testing"

	sys "github.com/docker-desktop/GoDockerInfoFetcher/pkg/sys"
)

func TestIsDockerInstalled(t *testing.T) {
	installed, err := sys.IsDockerInstalled()
	if err != nil {
		t.Errorf("IsDockerInstalled() error = %v, wantErr false", err)
		return
	}

	if !installed {
		t.Errorf("IsDockerInstalled() = %v, want true", installed)
	}
}

func TestGetDockerVersion(t *testing.T) {
	version := sys.GetDockerVersion()
	if version == "" {
		t.Errorf("GetDockerVersion() = %v, want not empty", version)
	}
}
