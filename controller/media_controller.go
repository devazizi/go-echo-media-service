package controller

import (
	"fmt"
	"github.com/devazizi/go-echo-media-service/infrastructure"
	"github.com/labstack/echo/v4"
	"net/http"
)

func UploadMedia(db infrastructure.DB) echo.HandlerFunc {

	return func(e echo.Context) error {

		file, _ := e.FormFile("file")
		collectionName := e.FormValue("collection_name")
		storageDisk := e.FormValue("storage_disk")

		fmt.Println(file, collectionName, storageDisk)

		return e.JSON(http.StatusCreated, map[string]any{
			"message": "hello echo framework",
		})
	}

}
