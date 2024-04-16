package main

import (
	"context"
	"fmt"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg/client"
)

func main() {
	// New Context
	ctx := context.Background()
	dockerSocket := "/var/run/docker.sock"
	client := client.NewClient(dockerSocket)

	// Get All Images
	images, err := client.ImageList(ctx)
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		// Print Image Info
		fmt.Println(image.ID)
	}
}
