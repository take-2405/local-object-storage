package domain

type CreateBucketRequest struct {
	Name string `json:"Name"`
}

type UploadImageRequest struct {
	Name   string `json:"Name"`
	Info   string `json:"Info"`
	Type   string `json:"Type"`
	Bucket string `json:"Bucket"`
}
