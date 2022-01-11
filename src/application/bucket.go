package application

type BucketUseCase interface {
	CreateBucket()
	GetBucketLists()
}

type bucketUseCase struct {
}

func NewBucketUseCase() BucketUseCase {
	return &bucketUseCase{}
}

func (bu bucketUseCase) CreateBucket() {

}

func (bu bucketUseCase) GetBucketLists() {

}
