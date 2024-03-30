package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"manntera.com/calculate-score-api/pkg/Database"
	"manntera.com/calculate-score-api/pkg/Usecase/CalculateBuffScoreFromImageUsecase"
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

type SkillResponse struct {
	Name       string `json:"name"`
	BaseParam  int    `json:"base_param"`
	BoostParam int    `json:"boost_param"`
}

type CalculateScoreResponse struct {
	Result       string          `json:"result"`
	ErrorMessage string          `json:"error_message"`
	Score        int             `json:"score"`
	Skills       []SkillResponse `json:"skills"`
}

func calculateScore(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse("invalid_form_data", err.Error()))
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, newErrorResponse("no_images", "No images uploaded"))
	}

	images, err := openImages(files)
	if err != nil {
		return c.JSON(http.StatusBadRequest, newErrorResponse("invalid_image", err.Error()))
	}
	defer closeImages(images)

	score, srcSkills, err := CalculateBuffScoreFromImageUsecase.CalculateBuffScoreFromImage(c.Request().Context(), images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse("internal_error", err.Error()))
	}

	skillResponse, err := buildSkillResponse(srcSkills)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse("internal_error", err.Error()))
	}

	response := CalculateScoreResponse{
		Result: "success",
		Score:  score,
		Skills: skillResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func newErrorResponse(errorCode, errorMessage string) CalculateScoreResponse {
	return CalculateScoreResponse{
		Result:       errorCode,
		ErrorMessage: errorMessage,
	}
}

func openImages(files []*multipart.FileHeader) ([]os.File, error) {
	var images []os.File
	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		tempFile, err := os.CreateTemp("", "image-*.png")
		if err != nil {
			return nil, err
		}

		_, err = io.Copy(tempFile, src)
		if err != nil {
			return nil, err
		}

		_, err = tempFile.Seek(0, 0)
		if err != nil {
			return nil, err
		}

		images = append(images, *tempFile)
	}
	return images, nil
}

func closeImages(images []os.File) {
	for _, image := range images {
		image.Close()
	}
}

func buildSkillResponse(srcSkills []Database.BuffSkillParam) ([]SkillResponse, error) {
	skillResponse := make([]SkillResponse, len(srcSkills))
	for i, srcSkill := range srcSkills {
		skill, err := Database.GetSkillFromId(srcSkill.SkillId)
		if err != nil {
			return nil, err
		}
		skillResponse[i] = SkillResponse{
			Name:       skill.Name,
			BaseParam:  srcSkill.BaseParam,
			BoostParam: srcSkill.BoostParam,
		}
	}
	return skillResponse, nil
}
