package types

type VersionInfo struct {
	Version string `json:"Version"`
	Os string `json:"Os"`
	KernelVersion string `json:"KernelVersion"`
	GoVersion string `json:"GoVersion"`
	GitCommit string `json:"GitCommit"`
	Arch string `json:"Arch"`
	BuildTime string `json:"BuildTime"`
	ApiVersion string `json:"ApiVersion"`
	Experimental bool `json:"Experimental"`
}

type Info struct {
	Architecture string `json:"Architecture"`
	ClusterStore string `json:"ClusterStore"`
	CgroupDriver string `json:"CgroupDriver"`
	Containers int `json:"Containers"`
	ContainersRunning int `json:"ContainersRunning"`
	ContainersStopped int `json:"ContainersStopped"`
	ContainersPaused int `json:"ContainersPaused"`
	CpuCfsPeriod bool `json:"CpuCfsPeriod"`
	CpuCfsQuota bool `json:"CpuCfsQuota"`
	Debug bool `json:"Debug"`
	DockerRootDir string `json:"DockerRootDir"`
	Driver string `json:"Driver"`
	DriverStatus [][]string `json:"DriverStatus"`
	ExperimentalBuild bool `json:"ExperimentalBuild"`
	HttpProxy string `json:"HttpProxy"`
	HttpsProxy string `json:"HttpsProxy"`
	ID string `json:"ID"`
	IPv4Forwarding bool `json:"IPv4Forwarding"`
	Images int `json:"Images"`
	IndexServerAddress string `json:"IndexServerAddress"`
	InitPath string `json:"InitPath"`
	InitSha1 string `json:"InitSha1"`
	KernelMemory bool `json:"KernelMemory"`
	KernelVersion string `json:"KernelVersion"`
	Labels []string `json:"Labels"`
	MemTotal int `json:"MemTotal"`
	MemoryLimit bool `json:"MemoryLimit"`
	NCPU int `json:"NCPU"`
	NEventsListener int `json:"NEventsListener"`
	NFd int `json:"NFd"`
	NGoroutines int `json:"NGoroutines"`
	Name string `json:"Name"`
	NoProxy string `json:"NoProxy"`
	OSType string `json:"OSType"`
	OomKillDisable bool `json:"OomKillDisable"`
	OperatingSystem string `json:"OperatingSystem"`
	RegistryConfig struct {
		AllowNondistributableArtifactsCIDRs []string `json:"AllowNondistributableArtifactsCIDRs"`
		AllowNondistributableArtifactsHostnames []string `json:"AllowNondistributableArtifactsHostnames"`
		InsecureRegistryCIDRs []string `json:"InsecureRegistryCIDRs"`
		IndexConfigs map[string]struct {
			Mirrors []string `json:"Mirrors"`
			Name string `json:"Name"`
			Official bool `json:"Official"`
			Secure bool `json:"Secure"`
		} `json:"IndexConfigs"`
		Mirrors []string `json:"Mirrors"`
		NonGlobalRegistryHostnames []string `json:"NonGlobalRegistryHostnames"`
		PermissiveHostnames []string `json:"PermissiveHostnames"`
	} `json:"RegistryConfig"`
	SecurityOptions []string `json:"SecurityOptions"`
	ServerVersion string `json:"ServerVersion"`
	SwapLimit bool `json:"SwapLimit"`
	SystemTime string `json:"SystemTime"`
}
