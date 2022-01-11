package application

type UploadUseCase interface {
	UploadImage()
}

type uploadUseCase struct {
}

func NewCreateBucketUseCase() UploadUseCase {
	return &uploadUseCase{}
}

func (uu uploadUseCase) UploadImage() {

}
