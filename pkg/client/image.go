package client

import (
	"context"
	"encoding/json"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

// Get All Images
func (cli *client) ImageList(ctx context.Context) ([]types.ImageSummary, error) {
	path := "/images/json?all=1"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var images []types.ImageSummary
	err = json.NewDecoder(resp.Body).Decode(&images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// Get Image By Image ID
func (cli *client) ImageInspect(ctx context.Context, imageID string) (types.ImageDetails, error) {
	path := "/images/" + imageID + "/json"

	resp, err := cli.get(ctx, path, map[string]string{"Content-Type": "application/json"})
	if err != nil {
		return types.ImageDetails{}, err
	}
	defer resp.Body.Close()

	var image types.ImageDetails
	err = json.NewDecoder(resp.Body).Decode(&image)
	if err != nil {
		return types.ImageDetails{}, err
	}

	return image, nil
}
