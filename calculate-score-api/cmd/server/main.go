package main

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"manntera.com/calculate-score-api/pkg/Database"
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

type ParamResponse struct {
	Id    int `json:"id"`
	Value int `json:"value"`
}
type SkillResponse struct {
	Name   string          `json:"name"`
	Params []ParamResponse `json:"params"`
}
type CalculateScoreResponse struct {
	Ok     bool            `json:"ok"`
	Error  string          `json:"error"`
	Score  int             `json:"score"`
	Skills []SkillResponse `json:"skills"`
}

func calculateScore(c echo.Context) error {
	form, formErr := c.MultipartForm()
	if formErr != nil {
		return c.JSON(http.StatusBadRequest, CalculateScoreResponse{
			Ok:    false,
			Error: "invalid_form_data",
		})
	}

	files := form.File["images"]
	if len(files) == 0 {
		return c.JSON(http.StatusBadRequest, CalculateScoreResponse{
			Ok:    false,
			Error: "no_images",
		})
	}

	images, readFilesErr := readFiles(files)

	if readFilesErr != nil {
		return c.JSON(http.StatusBadRequest, CalculateScoreResponse{
			Ok:    false,
			Error: "invalid_image",
		})
	}

	score, srcSkills, err := CalculateBuffScoreFromImageUsecase.CalculateBuffScoreFromImage(c.Request().Context(), images)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, CalculateScoreResponse{
			Ok:    false,
			Error: "internal_error",
		})
	}

	skillResponse := make([]SkillResponse, len(srcSkills))
	for i, skill := range srcSkills {
		paramResponse := make([]ParamResponse, len(skill.BuffParams))
		for j, param := range skill.BuffParams {
			paramResponse[j] = ParamResponse{
				Id:    param.ParamId,
				Value: int(param.ParamValue),
			}
		}
		srcSkill := Database.GetSkillFromId(skill.SkillId)
		skillResponse[i] = SkillResponse{
			Name:   srcSkill.Name,
			Params: paramResponse,
		}
	}
	response := CalculateScoreResponse{
		Ok:     true,
		Error:  "",
		Score:  score,
		Skills: skillResponse,
	}
	return c.JSON(http.StatusOK, response)
}

func readFiles(files []*multipart.FileHeader) ([]*os.File, error) {
	var images []*os.File
	for _, file := range files {
		src, srcErr := file.Open()
		if srcErr != nil {
			return nil, srcErr
		}
		defer src.Close()

		tempFile, tempErr := os.CreateTemp("", "image-*.png")
		if tempErr != nil {
			return nil, tempErr
		}
		defer tempFile.Close()

		_, copyErr := io.Copy(tempFile, src)
		if copyErr != nil {
			return nil, copyErr
		}

		_, seekErr := tempFile.Seek(0, 0)
		if seekErr != nil {
			return nil, seekErr
		}

		images = append(images, tempFile)
	}
	return images, nil
}
