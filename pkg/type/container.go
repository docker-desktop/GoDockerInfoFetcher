package types

type ContainerSummary struct {
	ID         string            `json:"Id"`
	Names      []string          `json:"Names"`
	Image      string            `json:"Image"`
	ImageID    string            `json:"ImageID"`
	Command    string            `json:"Command"`
	Created    int64             `json:"Created"`
	State      string            `json:"State"`
	Status     string            `json:"Status"`
	Ports      []Port            `json:"Ports"`
	Labels     map[string]string `json:"Labels"`
	SizeRw     int64             `json:"SizeRw,omitempty"`
	SizeRootFs int64             `json:"SizeRootFs,omitempty"`
}

type Port struct {
	PrivatePort int    `json:"PrivatePort"`
	PublicPort  int    `json:"PublicPort,omitempty"`
	Type        string `json:"Type"`
	IP          string `json:"IP,omitempty"`
}

type ContainerDetails struct {
	ID              string          `json:"Id"`
	Name            string          `json:"Name"`
	ImageID         string          `json:"Image"`               // Storing the image ID, renamed field for clarity
	ImageName       string          `json:"ImageName,omitempty"` // Additional field for the image name
	Created         string          `json:"Created"`
	Path            string          `json:"Path"`
	Args            []string        `json:"Args"`
	State           ContainerState  `json:"State"`
	ResolvConfPath  string          `json:"ResolvConfPath"`
	HostnamePath    string          `json:"HostnamePath"`
	HostsPath       string          `json:"HostsPath"`
	LogPath         string          `json:"LogPath"`
	Node            NodeDetails     `json:"Node,omitempty"`
	NetworkSettings NetworkSettings `json:"NetworkSettings"`
	Mounts          []Mount         `json:"Mounts"`
	Config          ContainerConfig `json:"Config"`
}

type ContainerState struct {
	Status     string `json:"Status"`
	Running    bool   `json:"Running"`
	Paused     bool   `json:"Paused"`
	Restarting bool   `json:"Restarting"`
	OOMKilled  bool   `json:"OOMKilled"`
	Dead       bool   `json:"Dead"`
	Pid        int    `json:"Pid"`
	ExitCode   int    `json:"ExitCode"`
	Error      string `json:"Error"`
	StartedAt  string `json:"StartedAt"`
	FinishedAt string `json:"FinishedAt"`
}

type NodeDetails struct {
}

type NetworkSettings struct {
	IPAddress   string `json:"IPAddress"`
	IPPrefixLen int    `json:"IPPrefixLen"`
	Gateway     string `json:"Gateway"`
	Bridge      string `json:"Bridge"`
}

type Mount struct {
	Source      string `json:"Source"`
	Destination string `json:"Destination"`
	Mode        string `json:"Mode"`
	RW          bool   `json:"RW"`
}

type ContainerConfig struct {
	Hostname     string              `json:"Hostname"`
	Domainname   string              `json:"Domainname"`
	User         string              `json:"User"`
	AttachStdin  bool                `json:"AttachStdin"`
	AttachStdout bool                `json:"AttachStdout"`
	AttachStderr bool                `json:"AttachStderr"`
	ExposedPorts map[string]struct{} `json:"ExposedPorts"`
	Tty          bool                `json:"Tty"`
	OpenStdin    bool                `json:"OpenStdin"`
	StdinOnce    bool                `json:"StdinOnce"`
	Env          []string            `json:"Env"`
	Cmd          []string            `json:"Cmd"`
}
