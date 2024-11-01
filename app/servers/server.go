package servers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StartServer() {
	app := echo.New()
	app.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	app.Logger.Fatal(app.Start(":1323"))
}
