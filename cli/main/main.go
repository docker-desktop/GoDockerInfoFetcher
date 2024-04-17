package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg/client"
)

func main() {
	// New Context
	ctx := context.Background()
	dockerSocket := "/var/run/docker.sock"
	client := client.NewClient(dockerSocket)

	containers, err := client.ContainerList(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	for _, container := range containers {
		log.Println(container.ID, container.Names, container.Image)
	}

	container, err := client.ContainerInspect(ctx, containers[0].ID)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(container.Name, container.ID)
}
