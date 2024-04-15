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

	log.Println("All Images")
	if images, err := client.ImageList(ctx); err != nil {
		panic(err)
	} else {
		for _, image := range images {
			fmt.Println(image.Containers)
		}
	}
}
