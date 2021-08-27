package usecases

import "github.com/polluxdev/clean-arch/app/domain"

type PostInteractor struct {
	PostRepository PostRepository
}

func (pi *PostInteractor) Index() (posts domain.Posts, err error) {
	posts, err = pi.PostRepository.FindAll()

	return
}
