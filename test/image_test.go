package test

import (
	"testing"

	GoDockerInfoFetcher "github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

var (
	dockerImageFetcher = GoDockerInfoFetcher.NewDockerImageFetcher()
)

func TestGetDockerImages(t *testing.T) {
	out, err := dockerImageFetcher.GetDockerImages()
	if err != nil {
		t.Errorf("GetDockerImages() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetDockerImages() returned empty list, expected at least one image")
	}
}
