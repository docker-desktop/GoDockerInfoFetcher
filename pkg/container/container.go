package godockerinfofetcher

import (
	"bytes"
	"os/exec"
	"strings"
)

type DockerContainer struct {
	ContainerID string
	Image       string
	Command     string
	Created     string
	Status      string
	Ports       string
	Names       string
}

func parseDockerPsOutput(output string) []DockerContainer {
	var containers []DockerContainer
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		// 탭으로 문자열을 분리
		parts := strings.Split(line, "\t")
		if len(parts) < 7 {
			continue
		}
		containers = append(containers, DockerContainer{
			ContainerID: parts[0],
			Image:       parts[1],
			Command:     parts[2],
			Created:     parts[3],
			Status:      parts[4],
			Ports:       parts[5],
			Names:       parts[6],
		})
	}
	return containers
}

// All
func GetDockerContainers() ([]DockerContainer, error) {
	cmd := exec.Command("docker", "ps", "-a", "--format", "{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}\t{{.Names}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return parseDockerPsOutput(out.String()), nil
}

// Only running containers
func GetRunningDockerContainers() ([]DockerContainer, error) {
	cmd := exec.Command("docker", "ps", "--format", "{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}\t{{.Names}}")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return parseDockerPsOutput(out.String()), nil
}
