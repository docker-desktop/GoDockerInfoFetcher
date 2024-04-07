package GoDockerInfoFetcher

import (
	"bytes"
	"os/exec"
	"strings"
)

// DockerContainer 정보를 나타내는 구조체
type DockerContainer struct {
	ContainerID string `json:"container_id"`
	Image       string `json:"image"`
	Command     string `json:"command"`
	Created     string `json:"created"`
	Status      string `json:"status"`
	Ports       string `json:"ports"`
	Names       string `json:"names"`
}

type DockerContainerFetcher interface {
	GetAllContainers() ([]DockerContainer, error)
	GetRunningContainers() ([]DockerContainer, error)
}

type dockerContainerFetcher struct{}

func NewDockerContainerFetcher() DockerContainerFetcher {
	return &dockerContainerFetcher{}
}

func (d dockerContainerFetcher) GetAllContainers() ([]DockerContainer, error) {
	return fetchContainers("docker", "ps", "-a", "--format", "{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}\t{{.Names}}")
}

func (d dockerContainerFetcher) GetRunningContainers() ([]DockerContainer, error) {
	return fetchContainers("docker", "ps", "--format", "{{.ID}}\t{{.Image}}\t{{.Command}}\t{{.CreatedAt}}\t{{.Status}}\t{{.Ports}}\t{{.Names}}")
}

func fetchContainers(command string, args ...string) ([]DockerContainer, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return parseDockerContainerPsOutput(out.String()), nil
}

func parseDockerContainerPsOutput(output string) []DockerContainer {
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
