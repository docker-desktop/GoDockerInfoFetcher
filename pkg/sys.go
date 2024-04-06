package GoDockerInfoFetcher

import (
	"errors"
	"os/exec"
)

// Docker가 설치되어 있는지 확인하는 함수
func IsDockerInstalled() (bool, error) {
	cmd := exec.Command("docker", "--version")
	if err := cmd.Run(); err != nil {
		return false, err
	}

	return true, nil
}

// Docker Version 정보를 가져오는 함수
func GetDockerVersion() string {
	cmd := exec.Command("docker", "--version")
	out, err := cmd.Output()
	if err != nil {
		return ""
	}

	return string(out)
}

// Docker Image가 하나라도 있는지 확인하는 함수
func IsDockerImageAvailable() (bool, error) {
	dockerImageFetcher := NewDockerImageFetcher()

	images, err := dockerImageFetcher.GetDockerImages()
	if err != nil {
		return false, err
	}

	if len(images) == 0 {
		return false, errors.New("no Docker images found on the system")
	}

	return true, nil
}

// Docker Container가 하나라도 있는지 확인하는 함수
func IsDockerContainerAvailable() (bool, error) {
	dockerContainerFetcher := NewDockerContainerFetcher()

	containers, err := dockerContainerFetcher.GetAllContainers()
	if err != nil {
		return false, err
	}

	if len(containers) == 0 {
		return false, errors.New("no Docker containers found on the system")
	}

	return true, nil
}

// Docker 실행중인 Container가 하나라도 있는지 확인하는 함수
func IsDockerRunningContainerAvailable() (bool, error) {
	dockerContainerFetcher := NewDockerContainerFetcher()

	containers, err := dockerContainerFetcher.GetRunningContainers()
	if err != nil {
		return false, err
	}

	if len(containers) == 0 {
		return false, errors.New("no Docker running containers found on the system")
	}

	return true, nil
}
