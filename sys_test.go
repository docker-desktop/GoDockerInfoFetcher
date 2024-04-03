package godockerinfofetcher

import "testing"

func TestIsDockerInstalled(t *testing.T) {
	installed := IsDockerInstalled()
	if !installed {
		t.Errorf("IsDockerInstalled() = %v, want true", installed)
	}
}

func TestGetDockerVersion(t *testing.T) {
	version := GetDockerVersion()
	if version == "" {
		t.Errorf("GetDockerVersion() = %v, want not empty", version)
	}
}
