package main

import (
	"log"
	"net/http"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/labstack/echo"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e := echo.New()
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{
	// 		"http://localhost:3000",
	// 	},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	// }))
	e.POST("/calculate-score", calculateScore)

	e.Logger.Fatal(e.Start(":" + port))
}

func calculateScore(c echo.Context) error {
	log.Default().Println("run calculate score")
	file, fileErr := c.FormFile("image")

	if fileErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "image is required",
		})
	}
	log.Default().Println("file:", file.Filename)

	client, visionErr := vision.NewImageAnnotatorClient(c.Request().Context())
	if visionErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to create vision client",
		})
	}
	log.Default().Println("create client")

	defer client.Close()

	src, srcErr := file.Open()

	if srcErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to open image",
		})
	}
	log.Default().Println("create src:")
	defer src.Close()

	image, imageErr := vision.NewImageFromReader(src)

	if imageErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to create image from reader",
		})
	}
	log.Default().Println("create image")

	annotations, detectErr := client.DetectTexts(c.Request().Context(), image, nil, 10)

	if detectErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to detect text",
		})
	}
	log.Default().Println("create annotations")

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
