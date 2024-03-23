package main

import (
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
	e.POST("/calculate-score", calculateScore)

	e.Logger.Fatal(e.Start(":" + port))
}

func calculateScore(c echo.Context) error {
	form, formErr := c.MultipartForm()
	if formErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "failed to parse multipart form",
		})
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "images are required",
		})
	}

	client, visionErr := vision.NewImageAnnotatorClient(c.Request().Context())
	if visionErr != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "failed to create vision client",
		})
	}
	defer client.Close()

	var results []map[string]interface{}

	for _, file := range files {
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

		var texts []string
		for _, annotation := range annotations {
			texts = append(texts, annotation.Description)
		}

		result := map[string]interface{}{
			"filename": file.Filename,
			"texts":    texts,
		}
		results = append(results, result)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"results": results,
	})
}
