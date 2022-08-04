package server

import (
	"goyo/libs"
	"goyo/modules/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitEcho() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	e.Logger.Fatal(e.Start(":" + libs.ENV.Port))
}

func InitGin() {
	server := gin.Default()
	server.Use(cors.Default())
	server.Use(middlewares.ErrorHandleRecovery())
	GinRoutes(server)
	server.Run(":8000")
}
