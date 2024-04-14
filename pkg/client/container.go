package client

import (
	"context"
	"encoding/json"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

func (cli *client) ContainerList(ctx context.Context) ([]types.Container, error) {
	path := "/containers/json"

	req, err := buildRequest(ctx, "GET", path)
	if err != nil {
		return nil, err
	}

	resp, err := cli.doRequest(req)
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
