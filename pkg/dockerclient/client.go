package dockerclient

import "github.com/docker/docker/client"

func Client() (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv)
}