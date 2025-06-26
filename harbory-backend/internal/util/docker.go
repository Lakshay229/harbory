package util

import (
	"log"

	"github.com/docker/docker/client"
)

var dockerClient *client.Client

func init() {
	var err error
	dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Failed to create Docker client: %v", err)
	}
}

func GetDockerClient() *client.Client {
	return dockerClient
}
