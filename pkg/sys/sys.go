package sys

import (
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
