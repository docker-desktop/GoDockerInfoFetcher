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

	fmt.Println("Contianer ID:", cnt_insp.ID)
	fmt.Println("Container Status:", cnt_insp.State.Status)

	err = client.ContainerStartByID(ctx, cnt_insp.ID) 
	if err != nil {
		log.Fatalln(err.Error()) 
	}
}
