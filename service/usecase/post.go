package usecase

type PostRepository interface {
	Create() error
	Read() error
	Update() error
	Delete() error
}

type PostUseCase struct {
	repository PostRepository
}

func NewPostUseCase(postRepository PostRepository) *PostUseCase {
	return &PostUseCase{
		repository: postRepository,
	}
}

func (u *PostUseCase) POstUser() error {
	err := u.repository.Create()
	if err != nil {
		return err
	}
	return nil
}

//TODO: Rest of post functions...
