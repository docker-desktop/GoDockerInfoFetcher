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

