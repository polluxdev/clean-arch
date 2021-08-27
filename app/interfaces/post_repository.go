package interfaces

import (
	"fmt"

	"github.com/polluxdev/clean-arch/app/domain"
)

type PostRepository struct {
	SQLHandler SQLHandler
}

func (pr *PostRepository) FindAll() (posts domain.Posts, err error) {
	post := domain.Post{}

	rows := pr.SQLHandler.Conn.Find(&post)

	fmt.Println(rows)
	return domain.Posts{}, nil
}
