package usecases

import "github.com/polluxdev/clean-arch/app/domain"

type PostRepository interface {
	FindAll() (domain.Posts, error)
}
