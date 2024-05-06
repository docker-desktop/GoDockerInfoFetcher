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
	defer client.Close()

	images, err := client.ImageList(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	img_spec, _ := client.ImageInspect(ctx, images[0].ID)

	fmt.Println(img_spec.RepoTags[0])

	err = client.ImageDeleteByName(ctx, img_spec.RepoTags[0])
	fmt.Println(err)
}
