package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/polluxdev/clean-arch/app/infrastructure/routes"
	"github.com/polluxdev/clean-arch/app/interfaces"
	"github.com/polluxdev/clean-arch/app/usecases"
)

func Dispatch(logger usecases.Logger, sqlHandler interfaces.SQLHandler) {
	postController := interfaces.NewPostController(sqlHandler, logger)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "메가야 사랑해요!")
	})

	v := e.Group("/api/v1")

	routes.PostRoutes(v, *postController)

	logger.LogAccess("%s", e.Start(":1323"))
}
