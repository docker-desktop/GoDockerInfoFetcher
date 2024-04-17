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

	containers, err := client.ContainerList(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Default().Println("All Containers")
	for _, container := range containers {
		log.Println(container.ID, container.Names, container.Image)
	}

	container, err := client.ContainerInspect(ctx, containers[0].ID)
	if err != nil {
		log.Fatalln(err)
	}
	log.Default().Println("Container Details")
	log.Println(container.ID, container.Name, container.ID)

	stoppedContainer, err := client.ContainerListStopped(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Default().Println("Stopped Containers")
	for _, container := range stoppedContainer {
		log.Println(container.ID, container.Names, container.Image)
	}

	runningContainer, err := client.ContainerListRunning(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Default().Println("Running Containers")
	for _, container := range runningContainer {
		log.Println(container.ID, container.Names, container.Image)
	}

}
