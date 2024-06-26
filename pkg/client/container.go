package client

import (
	"context"
	"encoding/json"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

// Get All Container List [docker ps -a]
func (cli *client) ContainerList(ctx context.Context) ([]types.ContainerSummary, error) {
	path := "/containers/json?all=1"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var containers []types.ContainerSummary
	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}

// Get Running Container List [docker ps]
func (cli *client) ContainerListRunning(ctx context.Context) ([]types.ContainerSummary, error) {
	path := "/containers/json?filters={\"status\":[\"running\"]}"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var containers []types.ContainerSummary
	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}

// Get Stopped Container List [docker ps -a]
func (cli *client) ContainerListStopped(ctx context.Context) ([]types.ContainerSummary, error) {
	path := "/containers/json?filters={\"status\":[\"exited\"]}"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, nil
	}
	defer resp.Body.Close()

	var containers []types.ContainerSummary
	err = json.NewDecoder(resp.Body).Decode(&containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}

// Get Container By Container ID
func (cli *client) ContainerInspect(ctx context.Context, containerID string) (types.ContainerDetails, error) {
	path := "/containers/" + containerID + "/json"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return types.ContainerDetails{}, err
	}
	defer resp.Body.Close()

	var container types.ContainerDetails
	err = json.NewDecoder(resp.Body).Decode(&container)
	if err != nil {
		return types.ContainerDetails{}, err
	}

	return container, nil
}

// Delete Container By Container ID
func (cli *client) ContainerDeleteByID(ctx context.Context, containerID string) error {
	path := "/containers/" + containerID

	resp, err := cli.delete(ctx, path, nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// Start Container By Container ID 
func (cli *client) ContainerStartByID(ctx context.Context, containerID string) error {
	path := "/containers/" + containerID + "/start"

	resp, err := cli.post(ctx, path, nil) 
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

// STOP Container By Container ID 
func (cli *client) ContainerStopByID(ctx context.Context, containerID string) error {
	path := "/containers/" + containerID + "/stop"

	resp, err := cli.post(ctx, path, nil) 
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
