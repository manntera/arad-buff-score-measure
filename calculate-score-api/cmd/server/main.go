package main

import (
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))
	e.POST("/calculate-score", calculateScore)

	e.Logger.Fatal(e.Start(":" + port))
}

func calculateScore(c echo.Context) error {
	file, fileErr := c.FormFile("image")

	if fileErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "image is required",
		})
	}

	client, visionErr := vision.NewImageAnnotatorClient(c.Request().Context())

	if visionErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to create vision client",
		})
	}

	defer client.Close()

	src, srcErr := file.Open()

	if srcErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to open image",
		})
	}

	defer src.Close()

	image, imageErr := vision.NewImageFromReader(src)

	if imageErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to create image from reader",
		})
	}

	annotations, detectErr := client.DetectTexts(c.Request().Context(), image, nil, 10)

	if detectErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to detect text",
		})
	}

	if len(annotations) == 0 {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"error": "no text found",
		})
	} else {

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Text": annotations[0].Description,
		})
	}
}
