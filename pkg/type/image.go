package types

type ImageSummary struct {
	ID          string            `json:"Id"`
	RepoTags    []string          `json:"RepoTags"`
	RepoDigests []string          `json:"RepoDigests"`
	Created     int64             `json:"Created"`
	Size        int64             `json:"Size"`
	VirtualSize int64             `json:"VirtualSize"`
	Labels      map[string]string `json:"Labels"`
	Containers  int               `json:"Containers"`
}

type ImageDetails struct {
	ID           string            `json:"Id"`
	ParentID     string            `json:"ParentId"`
	RepoTags     []string          `json:"RepoTags"`
	RepoDigests  []string          `json:"RepoDigests"`
	Created      string            `json:"Created"`
	Size         int64             `json:"Size"`
	VirtualSize  int64             `json:"VirtualSize"`
	Labels       map[string]string `json:"Labels"`
	Architecture string            `json:"Architecture"`
	Os           string            `json:"Os"`
	Config       struct {
		ExposedPorts map[string]struct{} `json:"ExposedPorts"`
		Env          []string            `json:"Env"`
		Cmd          []string            `json:"Cmd"`
		Volumes      map[string]struct{} `json:"Volumes"`
		WorkingDir   string              `json:"WorkingDir"`
		Labels       map[string]string   `json:"Labels"`
	}
	History []struct {
		Created   int64  `json:"Created"`
		CreatedBy string `json:"CreatedBy"`
		Comment   string `json:"Comment"`
	}
}
