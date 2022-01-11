package handler

type BucketListResponse struct {
	Buckets []string `json:"Buckets"`
}

func ReturnBucketListResponse(Bukets []string) BucketListResponse {
	body := BucketListResponse{
		Buckets: Bukets,
	}
	return body
}

type Error struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func ReturnErrorResponse(code int, msg, desc string) Error {
	body := Error{
		Code:        code,
		Message:     msg,
		Description: desc,
	}
	return body
}
