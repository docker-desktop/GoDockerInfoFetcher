package types

type Image struct {
	ID          string            `json:"Id"`          // 이미지의 ID
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
