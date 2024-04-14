package client

import (
	"context"
	"encoding/json"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

// Get All Container List [docker ps -a]
func (cli *client) ContainerList(ctx context.Context) ([]types.Container, error) {
	path := "/containers/json?all=1"

	resp, err := cli.get(ctx, path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var containers []types.Container
	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}

// Get Running Container List [docker ps]
func (cli *client) ContainerListRunning(ctx context.Context) ([]types.Container, error) {
	path := "/containers/json?filters={\"status\":[\"running\"]}"

	resp, err := cli.get(ctx, path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var containers []types.Container
	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}
