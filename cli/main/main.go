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
	defer client.Close()

	containers, err := client.ContainerList(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(containers)
}
