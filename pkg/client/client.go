package client

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/http"
	"encoding/json"
	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

type Client interface {
	GetInfo(ctx context.Context) (types.Info, error)
	GetVersion(ctx context.Context) (types.VersionInfo, error)
	Close() error

	ContainerList(ctx context.Context) ([]types.ContainerSummary, error)
	ContainerListRunning(ctx context.Context) ([]types.ContainerSummary, error)
	ContainerListStopped(ctx context.Context) ([]types.ContainerSummary, error)
	ContainerInspect(ctx context.Context, containerID string) (types.ContainerDetails, error)
	ContainerDeleteByID(ctx context.Context, containerID string) error
	ContainerStartByID(ctx context.Context, containerID string) error
	ContainerStopByID(ctx context.Context, containerID string) error

	ImageList(ctx context.Context) ([]types.ImageSummary, error)
	ImageInspect(ctx context.Context, imageID string) (types.ImageDetails, error)
	ImageDeleteByName(ctx context.Context, name string) error
}

type client struct {
	Http *http.Client
	Conn *net.Conn
}

func newHTTPClient() *http.Client {
	transport := &http.Transport{}

	return &http.Client{
		Transport: transport,
	}
}

func newClient(dockerSock string) *net.Conn {
	conn, err := net.Dial("unix", dockerSock)
	if err != nil {
		log.Fatal("Error connecting to Docker socket:", err)
	}

	return &conn
}

func buildRequest(ctx context.Context, method, path string, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, path, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

func NewClient(dockerSock string) Client {
	conn := newClient(dockerSock)
	http_cli := newHTTPClient()

	return &client{
		Http: http_cli,
		Conn: conn,
	}
}

func (cli *client) doRequest(req *http.Request) (*http.Response, error) {
	err := req.Write(*cli.Conn)
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(*cli.Conn)
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *client) get(ctx context.Context, path string, headers map[string]string) (*http.Response, error) {
	req, err := buildRequest(ctx, http.MethodGet, path, headers)
	if err != nil {
		return nil, err
	}

	resp, err := cli.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *client) delete(ctx context.Context, path string, headers map[string]string) (*http.Response, error) {
	req, err := buildRequest(ctx, http.MethodDelete, path, headers)
	if err != nil {
		return nil, err
	}

	resp, err := cli.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *client) post(ctx context.Context, path string, headers map[string]string) (*http.Response, error) {
	req, err := buildRequest(ctx, http.MethodPost, path, headers)
	if err != nil {
		return nil, err
	}

	resp, err := cli.doRequest(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (cli *client) GetInfo(ctx context.Context) (types.Info, error) {
	path := "/info"

	resp, err := cli.get(ctx, path, nil)
	if err != nil {
		return types.Info{}, err
	}
	defer resp.Body.Close()

	var info types.Info	
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return types.Info{}, err
	}

	return info, nil
}

func (cli *client) GetVersion(ctx context.Context) (types.VersionInfo, error) {
	path := "/version"

	resp, err := cli.get(ctx, path, nil)
	if err != nil {
		return types.VersionInfo{}, err
	}
	defer resp.Body.Close()

	var version types.VersionInfo
	err = json.NewDecoder(resp.Body).Decode(&version)
	if err != nil {
		return types.VersionInfo{}, err
	}

	return version, nil
}

func (cli *client) Close() error {
	return (*cli.Conn).Close()
}

