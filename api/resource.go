package api

type Resource struct {
	ID int `json:"id"`

	// Standard fields
	CreatorID int   `json:"creatorId"`
	CreatedTs int64 `json:"createdTs"`
	UpdatedTs int64 `json:"updatedTs"`

	// Domain specific fields
	Filename     string     `json:"filename"`
	Blob         []byte     `json:"-"`
	InternalPath string     `json:"internalPath"`
	ExternalLink string     `json:"externalLink"`
	Type         string     `json:"type"`
	Size         int64      `json:"size"`
	Visibility   Visibility `json:"visibility"`

	// Related fields
	LinkedMemoAmount int `json:"linkedMemoAmount"`
}

type ResourceCreate struct {
	// Standard fields
	CreatorID int `json:"-"`

	// Domain specific fields
	Filename     string     `json:"filename"`
	Blob         []byte     `json:"-"`
	InternalPath string     `json:"internalPath"`
	ExternalLink string     `json:"externalLink"`
	Type         string     `json:"type"`
	Size         int64      `json:"-"`
	Visibility   Visibility `json:"visibility"`
}

type ResourceFind struct {
	ID *int `json:"id"`

	// Standard fields
	CreatorID *int `json:"creatorId"`

	// Domain specific fields
	Filename *string `json:"filename"`
	MemoID   *int
	GetBlob  bool
}

type ResourcePatch struct {
	ID int `json:"-"`

	// Standard fields
	UpdatedTs *int64

	// Domain specific fields
	Filename   *string     `json:"filename"`
	Visibility *Visibility `json:"visibility"`
}

type ResourceDelete struct {
	ID int
}
