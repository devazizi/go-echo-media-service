package main

import (
	"github.com/devazizi/go-echo-media-service/controller"
	"github.com/devazizi/go-echo-media-service/infrastructure"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")

	dsn := "root:@tcp(127.0.0.1:3306)/go_media?charset=utf8mb4&parseTime=True&loc=Local"
	database := infrastructure.NewDB(dsn)

	e := echo.New()

	v1 := e.Group("api/v1")
	{
		v1.POST("/media", controller.UploadMedia(database))
	}

	e.Logger.Fatal(e.Start("0.0.0.0:5000"))
}
