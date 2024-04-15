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

type Image struct {
	Id          string            `json:"Id"`          // 이미지의 ID
	ParentID    string            `json:"ParentId"`    // 부모 이미지의 ID
	RepoTags    []string          `json:"RepoTags"`    // 이미지의 리포지토리 태그 목록
	RepoDigests []string          `json:"RepoDigests"` // 이미지의 리포지토리 다이제스트 목록
	Created     int64             `json:"Created"`     // 이미지가 생성된 시간 (Unix 타임스탬프)
	Size        int64             `json:"Size"`        // 이미지의 크기 (바이트 단위)
	SharedSize  int64             `json:"SharedSize"`  // 공유된 이미지의 크기
	VirtualSize int64             `json:"VirtualSize"` // 이미지의 전체 가상 크기
	Labels      map[string]string `json:"Labels"`      // 이미지에 설정된 라벨 (키-값 쌍)
	Containers  int64             `json:"Containers"`  // 이미지를 사용하는 컨테이너 수
}
