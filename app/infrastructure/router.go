package infrastructure

import (
	"net/http"

	"github.com/labstack/echo"
)

func Dispatch() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "메가야 사랑해요!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
