package client

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/http"

	types "github.com/docker-desktop/GoDockerInfoFetcher/pkg/type"
)

type Client interface {
	Close() error
	ContainerList(ctx context.Context) ([]types.Container, error)
	ContainerListRunning(ctx context.Context) ([]types.Container, error)
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

func (cli *client) Close() error {
	return (*cli.Conn).Close()
}
