package main

import (
	"context"
	"log"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg/client"
)

func main() {
	// New Context
	ctx := context.Background()
	dockerSocket := "/var/run/docker.sock"
	client := client.NewClient(dockerSocket)

	log.Println("All Containers")
	if containers, err := client.ContainerList(ctx); err != nil {
		panic(err)
	} else {
		for _, container := range containers {
			println(container.Id, container.Image, container.Status)
		}
	}

	log.Println("Running Containers")
	if containers, err := client.ContainerListRunning(ctx); err != nil {
		panic(err)
	} else {
		for _, container := range containers {
			println(container.Id, container.Image, container.Status)
		}
	}
}
