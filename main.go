package main

import (
	"context"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg/client"
)

func main() {
	// New Context
	ctx := context.Background()
	dockerSocket := "/var/run/docker.sock"
	client := client.NewClient(dockerSocket)

	containers, err := client.ContainerList(ctx)
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		println(container.Id, container.Image, container.Status)
	}

}
