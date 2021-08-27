package interfaces

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/polluxdev/clean-arch/app/usecases"
)

type PostController struct {
	PostInteractor usecases.PostInteractor
	Logger         usecases.Logger
}

func NewPostController(sqlHandler SQLHandler, logger usecases.Logger) *PostController {
	return &PostController{
		PostInteractor: usecases.PostInteractor{
			PostRepository: &PostRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

func (pc *PostController) Index(c echo.Context) error {
	posts, err := pc.PostInteractor.Index()
	if err != nil {
		pc.Logger.LogError("%s", err)
		c.String(http.StatusBadRequest, "ERROR")
	}
	fmt.Println(posts)
	return c.String(http.StatusOK, "Post List")
}
