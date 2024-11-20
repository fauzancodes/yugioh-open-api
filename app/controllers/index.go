package controllers

import (
	"bytes"
	"context"
	"io"
	"net/http"

	"github.com/Backblaze/blazer/b2"
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/dto"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	keyID := config.LoadConfig().BackblazeKeyID
	applicationKey := config.LoadConfig().BackblazeApplicationKey
	bucketName := config.LoadConfig().BackblazeBucketName
	folder := config.LoadConfig().BackblazeFolder
	ctx := context.Background()

	b2, err := b2.NewClient(ctx, keyID, applicationKey)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to connect to Backblaze",
				Error:   err.Error(),
			},
		)
	}

	bucket, err := b2.Bucket(ctx, bucketName)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.Response{
				Status:  http.StatusNotFound,
				Message: "Backblaze bucket not found",
				Error:   err.Error(),
			},
		)
	}

	reader := bucket.Object(folder + "/assets/html/index.html").NewReader(ctx)

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to read file",
				Error:   err.Error(),
			},
		)
	}

	return c.Blob(http.StatusOK, "text/html", buf.Bytes())
}

func DownloadPostmanCollection(c echo.Context) error {
	keyID := config.LoadConfig().BackblazeKeyID
	applicationKey := config.LoadConfig().BackblazeApplicationKey
	bucketName := config.LoadConfig().BackblazeBucketName
	folder := config.LoadConfig().BackblazeFolder
	ctx := context.Background()

	b2, err := b2.NewClient(ctx, keyID, applicationKey)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to connect to Backblaze",
				Error:   err.Error(),
			},
		)
	}

	bucket, err := b2.Bucket(ctx, bucketName)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.Response{
				Status:  http.StatusNotFound,
				Message: "Backblaze bucket not found",
				Error:   err.Error(),
			},
		)
	}

	reader := bucket.Object(folder + "/docs/Yu-Gi-Oh! Open API.postman_collection.json").NewReader(ctx)

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to read file",
				Error:   err.Error(),
			},
		)
	}

	c.Response().Header().Set("Content-Disposition", `attachment; filename="Yu-Gi-Oh! Open API.postman_collection.json"`)
	c.Response().Header().Set("Content-Type", "application/octet-stream")
	return c.Blob(http.StatusOK, "application/octet-stream", buf.Bytes())
}

func DownloadPostmanEnvironment(c echo.Context) error {
	keyID := config.LoadConfig().BackblazeKeyID
	applicationKey := config.LoadConfig().BackblazeApplicationKey
	bucketName := config.LoadConfig().BackblazeBucketName
	folder := config.LoadConfig().BackblazeFolder
	ctx := context.Background()

	b2, err := b2.NewClient(ctx, keyID, applicationKey)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to connect to Backblaze",
				Error:   err.Error(),
			},
		)
	}

	bucket, err := b2.Bucket(ctx, bucketName)
	if err != nil {
		return c.JSON(
			http.StatusNotFound,
			dto.Response{
				Status:  http.StatusNotFound,
				Message: "Backblaze bucket not found",
				Error:   err.Error(),
			},
		)
	}

	reader := bucket.Object(folder + "/docs/Yu-Gi-Oh! Open API.postman_environment.json").NewReader(ctx)

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, reader); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			dto.Response{
				Status:  http.StatusInternalServerError,
				Message: "Failed to read file",
				Error:   err.Error(),
			},
		)
	}

	c.Response().Header().Set("Content-Disposition", `attachment; filename="Yu-Gi-Oh! Open API.postman_environment.json"`)
	c.Response().Header().Set("Content-Type", "application/octet-stream")
	return c.Blob(http.StatusOK, "application/octet-stream", buf.Bytes())
}
