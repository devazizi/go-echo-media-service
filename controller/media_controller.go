package controller

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/devazizi/go-echo-media-service/entity"
	"github.com/devazizi/go-echo-media-service/infrastructure"
	"github.com/devazizi/go-echo-media-service/repository"
	"github.com/devazizi/go-echo-media-service/service/helpers"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
)

func UploadMedia(db infrastructure.DB) echo.HandlerFunc {

	return func(e echo.Context) error {

		file, err := e.FormFile("file")
		if err != nil {
			return err
		}

		collectionName := e.FormValue("collection_name")
		storageDisk := e.FormValue("storage_disk")

		region := os.Getenv("ARVAN_CLOUD_DEFAULT_REGION")
		endpoint := os.Getenv("ARVAN_CLOUD_ENDPOINT_URL")

		sess, err := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(
				os.Getenv("ARVAN_CLOUD_ACCESS_KEY"),
				os.Getenv("ARVAN_CLOUD_SECRET_KEY"),
				"",
			),
			Region:   aws.String(region),
			Endpoint: aws.String(endpoint),
		})

		media, err := file.Open()
		if err != nil {
			return err
		}
		defer media.Close()

		randomFileName := helpers.RandomString(100)
		uploader := s3manager.NewUploader(sess)
		uploadInfo, _ := uploader.Upload(&s3manager.UploadInput{

			Bucket: aws.String(os.Getenv("ARVAN_CLOUD_BUCKET")),

			Key: aws.String(randomFileName),

			Body: media,
		})

		uploadedMediaEntity := repository.NewDB(db).StoreMedia(entity.Media{
			StorageDisk: storageDisk,
			Name:        file.Filename,
			FileName:    randomFileName,
			Collection:  collectionName,
			MimeType:    file.Header.Get("Content-Type"),
			Url:         uploadInfo.Location,
		})

		return e.JSON(http.StatusCreated, map[string]any{
			"message": "upload file successfully",
			"media":   uploadedMediaEntity,
		})
	}

}
