package test

import (
	"testing"

	img "github.com/docker-desktop/GoDockerInfoFetcher/pkg/image"
)

func TestGetDockerImages(t *testing.T) {
	out, err := img.GetDockerImages()
	if err != nil {
		t.Errorf("GetDockerImages() error = %v, wantErr false", err)
		return
	}

	if len(out) == 0 {
		t.Error("GetDockerImages() returned empty list, expected at least one image")
	}
}
