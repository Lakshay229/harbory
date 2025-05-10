package main

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	if len(containers) == 0 {
		fmt.Println("No running containers found")
	} else {
		for _, container := range containers {
			fmt.Printf("Container ID: %s, Image: %s, Status: %s\n", container.ID, container.Image, container.Status)
		}
	}

	fmt.Println("Hello, Docker client initialized successfully")
}
