package types

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
