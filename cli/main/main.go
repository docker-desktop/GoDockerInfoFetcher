package main

import (
	"context"
	"fmt"
	"log"

	"github.com/docker-desktop/GoDockerInfoFetcher/pkg/client"
)

// 어떤 Image를 사용하여 Container가 만들어 졌다면 Image 삭제가 먼저 되어야 함
// Container가 Start 상태라면 Stop 후 삭제가 되어야 함

func main() {
	// New Context
	ctx := context.Background()
	dockerSocket := "/var/run/docker.sock"
	client := client.NewClient(dockerSocket)
	defer client.Close()

	containers, err := client.ContainerListStopped(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	cnt_insp, err := client.ContainerInspect(ctx, containers[0].ID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println(containers[0].ID)

	err = client.ContainerDeleteByID(ctx, cnt_insp.ID)
	if err != nil {
		log.Fatalln(err.Error())
	}

	images, err := client.ImageList(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	img_spec, _ := client.ImageInspect(ctx, images[0].ID)

	fmt.Println(img_spec.RepoTags[0])

	err = client.ImageDeleteByName(ctx, img_spec.RepoTags[0])
	if err != nil {
		log.Fatalln(err.Error())
	}
}
