package test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	GoDockerInfoFetcher "github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

// MockDockerImageFetcher는 DockerImageFetcher 인터페이스의 모의 구현
type MockDockerImageFetcher struct {
	mock.Mock
}

func (m *MockDockerImageFetcher) GetDockerImages() ([]GoDockerInfoFetcher.DockerImage, error) {
	args := m.Called()
	return args.Get(0).([]GoDockerInfoFetcher.DockerImage), args.Error(1)
}

func TestGetDockerImages(t *testing.T) {
	mockFetcher := new(MockDockerImageFetcher)
	dockerImages := []GoDockerInfoFetcher.DockerImage{
		{Repository: "busybox", Tag: "latest", ImageID: "abc123", Created: "2 days ago", Size: "1.23MB"},
	}
	mockFetcher.On("GetDockerImages").Return(dockerImages, nil)

	out, err := mockFetcher.GetDockerImages()
	if err != nil {
		t.Errorf("GetDockerImages() error = %v, wantErr false", err)
	}
	if len(out) == 0 {
		t.Error("GetDockerImages() returned empty list, expected at least one image")
	} else if out[0].Repository != "busybox" {
		t.Errorf("Expected 'busybox', got %s", out[0].Repository)
	}

	mockFetcher.AssertExpectations(t)
}
