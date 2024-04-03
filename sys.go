// Check Docker installed on PC OR Any other Programs

package godockerinfofetcher

import (
	"os/exec"
)

// Docker가 설치되어 있는지 확인하는 함수
func IsDockerInstalled() bool {
	cmd := exec.Command("docker", "--version")

	if err := cmd.Run(); err != nil {
		return false
	}
	return true
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
