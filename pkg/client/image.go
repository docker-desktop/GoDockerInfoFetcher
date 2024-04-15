package client

import (
	"context"
	"encoding/json"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

// Get All Images
func (cli *client) ImageList(ctx context.Context) ([]types.Image, error) {
	path := "/images/json"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var images []types.Image
	err = json.NewDecoder(resp.Body).Decode(&images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// Get Image By Image Name

// Get Image By Image ID

// Get Image By Image Tag

// Get Image By Image Repository

// Get Image By Image Created Time

// Get Image By Image Size

// Get Image By Image Virtual Size

// Get Image By Image Parent ID
