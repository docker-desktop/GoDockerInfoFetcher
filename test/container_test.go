package test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	GoDockerInfoFetcher "github.com/docker-desktop/GoDockerInfoFetcher/pkg"
)

// MockDockerContainerFetcher는 DockerContainerFetcher 인터페이스의 모의 구현
type MockDockerContainerFetcher struct {
	mock.Mock
}

func (m *MockDockerContainerFetcher) GetAllContainers() ([]GoDockerInfoFetcher.DockerContainer, error) {
	args := m.Called()
	return args.Get(0).([]GoDockerInfoFetcher.DockerContainer), args.Error(1)
}

func (m *MockDockerContainerFetcher) GetRunningContainers() ([]GoDockerInfoFetcher.DockerContainer, error) {
	args := m.Called()
	return args.Get(0).([]GoDockerInfoFetcher.DockerContainer), args.Error(1)
}

func TestGetAllContainers(t *testing.T) {
	mockFetcher := new(MockDockerContainerFetcher)
	dockerContainers := []GoDockerInfoFetcher.DockerContainer{
		{ContainerID: "1", Image: "busybox"},
	}
	mockFetcher.On("GetAllContainers").Return(dockerContainers, nil)

	out, err := mockFetcher.GetAllContainers()
	if err != nil {
		t.Errorf("GetAllContainers() error = %v, wantErr false", err)
	}
	if len(out) == 0 {
		t.Error("GetAllContainers() returned empty list, expected at least one container")
	}

	mockFetcher.AssertExpectations(t)
}

func TestGetRunningContainers(t *testing.T) {
	mockFetcher := new(MockDockerContainerFetcher)
	dockerContainers := []GoDockerInfoFetcher.DockerContainer{
		{ContainerID: "1", Image: "busybox", Status: "Up 10 minutes"},
	}
	mockFetcher.On("GetRunningContainers").Return(dockerContainers, nil)

	out, err := mockFetcher.GetRunningContainers()
	if err != nil {
		t.Errorf("GetRunningContainers() error = %v, wantErr false", err)
	}
	if len(out) == 0 {
		t.Error("GetRunningContainers() returned empty list, expected at least one running container")
	}

	mockFetcher.AssertExpectations(t)
}
