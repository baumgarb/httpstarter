package httpstarter

type ChangeTracking struct {
	LastModifiedBy string `json:"lastModifiedBy"`
	LastModifiedAt int    `json:"lastModifiedAt"`
}
