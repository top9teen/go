package route

import (
	"config"
	"controllers"

	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	// Routes
	g := e.Group("/api/employee")
	{
		g.GET("/load", controllers.GetAll)
		g.GET("/load/:count", controllers.Get)
	}

	// Routes for test
	t := e.Group("/api/test")
	{
		t.GET("/hello", controllers.HelloWorld)
	}

	e.Logger.Fatal(e.Start(config.SERVERPORT))

	return e
}
