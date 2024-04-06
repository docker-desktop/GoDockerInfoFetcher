package GoDockerInfoFetcher

import (
	"bytes"
	"os/exec"
	"strings"
)

type DockerImage struct {
	Repository string
	Tag        string
	ImageID    string
	Created    string
	Size       string
}

type DockerImageFetcher interface {
	GetDockerImages() ([]DockerImage, error)
}

type dockerImageFetcher struct{}

func NewDockerImageFetcher() DockerImageFetcher {
	return &dockerImageFetcher{}
}

func (d dockerImageFetcher) GetDockerImages() ([]DockerImage, error) {
	return fetchImages("docker", "images", "--format", "{{.Repository}}\t{{.Tag}}\t{{.ID}}\t{{.CreatedSince}}\t{{.Size}}")
}

func fetchImages(command string, args ...string) ([]DockerImage, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return parseDockerImagesOutput(out.String()), nil
}

func parseDockerImagesOutput(output string) []DockerImage {
	var images []DockerImage
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\t")
		if len(parts) < 5 {
			continue
		}
		images = append(images, DockerImage{
			Repository: parts[0],
			Tag:        parts[1],
			ImageID:    parts[2],
			Created:    parts[3],
			Size:       parts[4],
		})
	}
	return images
}
