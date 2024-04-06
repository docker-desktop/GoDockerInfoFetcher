package test

import (
	"fmt"
	"testing"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

var (
	dockerImageFetcher = pkg.NewDockerImageFetcher()
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

	fmt.Println(out)
}
