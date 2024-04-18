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

	imags, err := client.ImageList(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, img := range imags {
		fmt.Println("ID:", img.ID)
		fmt.Println("RepoTags:", img.RepoTags)
		fmt.Println("RepoDigests:", img.RepoDigests)
		fmt.Println("Created:", img.Created)
		fmt.Println("Size:", img.Size)
		fmt.Println("VirtualSize:", img.VirtualSize)
		fmt.Println("Labels:", img.Labels)
		fmt.Println("Containers:", img.Containers)
		fmt.Println()
	}

	inspectImage, err := client.ImageInspect(ctx, imags[0].ID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("ID:", inspectImage.ID)
	fmt.Println(inspectImage.Created, inspectImage.RepoTags)
}
