package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
)

type Port struct {
	IP          string `json:"IP"`
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort"`
	Type        string `json:"Type"`
}

type Container struct {
	Id              string            `json:"Id"`
	Names           []string          `json:"Names"`
	Image           string            `json:"Image"`
	ImageID         string            `json:"ImageID"`
	Command         string            `json:"Command"`
	Created         int64             `json:"Created"`
	Ports           []Port            `json:"Ports"`
	Labels          map[string]string `json:"Labels"`
	State           string            `json:"State"`
	Status          string            `json:"Status"`
	HostConfig      map[string]string `json:"HostConfig"`
	NetworkSettings NetworkSettings   `json:"NetworkSettings"`
	Mounts          []interface{}     `json:"Mounts"`
}

type NetworkSettings struct {
	Networks map[string]Network `json:"Networks"`
}

type Network struct {
	IPAMConfig          interface{} `json:"IPAMConfig"`
	Links               interface{} `json:"Links"`
	Aliases             interface{} `json:"Aliases"`
	MacAddress          string      `json:"MacAddress"`
	NetworkID           string      `json:"NetworkID"`
	EndpointID          string      `json:"EndpointID"`
	Gateway             string      `json:"Gateway"`
	IPAddress           string      `json:"IPAddress"`
	IPPrefixLen         int         `json:"IPPrefixLen"`
	IPv6Gateway         string      `json:"IPv6Gateway"`
	GlobalIPv6Address   string      `json:"GlobalIPv6Address"`
	GlobalIPv6PrefixLen int         `json:"GlobalIPv6PrefixLen"`
	DriverOpts          interface{} `json:"DriverOpts"`
	DNSNames            interface{} `json:"DNSNames"`
}

func main() {
	dockerSocket := "/var/run/docker.sock"
	path := "/containers/json"

	conn, err := net.Dial("unix", dockerSocket)
	if err != nil {
		log.Fatal("Error connecting to Docker socket:", err)
	}
	defer conn.Close()

	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Fatal("Error creating request:", err)
	}

	err = req.Write(conn)
	if err != nil {
		log.Fatal("Error writing request to Docker socket:", err)
	}

	reader := bufio.NewReader(conn)
	resp, err := http.ReadResponse(reader, req)
	if err != nil {
		log.Fatal("Error reading response:", err)
	}
	defer resp.Body.Close()

	var jsonData string
	buf := make([]byte, 1024)
	for {
		n, err := resp.Body.Read(buf)
		if n > 0 {
			jsonData += string(buf[:n])
		}
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal("Error reading response body:", err)
		}
	}

	var containers []Container
	err = json.Unmarshal([]byte(jsonData), &containers)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	for _, container := range containers {
		fmt.Printf("Container ID: %s, Image: %s, State: %s\n", container.Id, container.Image, container.State)
	}
}
