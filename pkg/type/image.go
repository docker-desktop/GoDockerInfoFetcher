package types

type ImageSummary struct {
	ID          string            `json:"Id"`
	ParentID    string            `json:"ParentId,omitempty"`
	RepoTags    []string          `json:"RepoTags"`
	RepoDigests []string          `json:"RepoDigests"`
	Created     int64             `json:"Created"` // Unix timestamp of when the image was created
	Size        int64             `json:"Size"`
	VirtualSize int64             `json:"VirtualSize"`
	Labels      map[string]string `json:"Labels,omitempty"`
	Containers  int               `json:"Containers"`
	SharedSize  int64             `json:"SharedSize,omitempty"` // Size of shared data
}
