package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
	CalculateBuffScoreFromImageUsecase "manntera.com/calculate-score-api/pkg/Usecase/CalculateBuffScoreFromImageUsecase"
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

	var images []os.File
	for _, file := range files {
		src, srcErr := file.Open()
		if srcErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "failed to open image",
			})
		}
		defer src.Close()

		tempFile, tempErr := os.CreateTemp("", "image-*.jpg")
		if tempErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "failed to create temporary file",
			})
		}
		defer tempFile.Close()

		_, copyErr := io.Copy(tempFile, src)
		if copyErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "failed to copy image data",
			})
		}

		// Seek to the beginning of the temporary file
		_, seekErr := tempFile.Seek(0, 0)
		if seekErr != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": "failed to seek temporary file",
			})
		}

		images = append(images, *tempFile)
	}

	score, err := CalculateBuffScoreFromImageUsecase.CalculateBuffScoreFromImage(c.Request().Context(), images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"score": score,
	})
}
