package image

import (
	"bufio"
	"bytes"
	"os/exec"
	"strings"
)

// DockerImage는 Docker 이미지의 정보를 나타냅니다.
type DockerImage struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
}

// parseDockerImagesOutput 함수는 `docker images` 명령의 출력을 분석
func parseDockerImagesOutput(output string) []DockerImage {
	images := []DockerImage{}
	scanner := bufio.NewScanner(strings.NewReader(output))
	scanner.Scan() // 첫 번째 라인(헤더)는 건너뜁니다.
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}
		images = append(images, DockerImage{
			Repository: fields[0],
			Tag:        fields[1],
			ImageID:    fields[2],
			Created:    strings.Join(fields[3:len(fields)-1], " "),
			Size:       fields[len(fields)-1],
		})
	}

	return images
}

// GetDockerImages 함수는 설치된 Docker 이미지 목록을 반환합니다.
func GetDockerImages() ([]DockerImage, error) {
	cmd := exec.Command("docker", "images")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}

	return parseDockerImagesOutput(out.String()), nil
}
