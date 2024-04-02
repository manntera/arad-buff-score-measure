package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"manntera.com/calculate-score-api/pkg/Repository/BuffEffectRepo"
	"manntera.com/calculate-score-api/pkg/Repository/BuffSkillRepo"
	"manntera.com/calculate-score-api/pkg/Repository/DetectedTextRepo"
	"manntera.com/calculate-score-api/pkg/Repository/SamplerImageRepo"
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

	filesHeaders := form.File["images"]
	if len(filesHeaders) == 0 {
		return c.JSON(http.StatusBadRequest, newErrorResponse("no_images", "No images uploaded"))
	}
	buffSkillJson, err := os.Getwd()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse("internal_error", err.Error()))
	}
	buffSkillJson += "/setting/BuffSkill.json"

	buffSkillRepo, err := BuffSkillRepo.NewBuffSkillRepoFromJsonFile(buffSkillJson)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrorResponse("internal_error", err.Error()))
	}
	buffEffectRepos := make([]*BuffEffectRepo.BuffEffectRepo, 0)
	for _, fileHeader := range filesHeaders {
		samplerImageRepo, err := SamplerImageRepo.NewSamplerImageRepoFromFileHeader(fileHeader)
		if err != nil {
			return c.JSON(http.StatusBadRequest, newErrorResponse("invalid_image", err.Error()))
		}
		defer samplerImageRepo.Close()
		detectedTextRepo, err := DetectedTextRepo.NewDetectedTextRepoFromSamplerImageRepo(c.Request().Context(), samplerImageRepo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, newErrorResponse("invalid_image", err.Error()))
		}
		buffEffectRepo, err := BuffEffectRepo.NewBuffEffectRepoFromDetectedTextRepo(buffSkillRepo, detectedTextRepo)
		if err != nil {
			return c.JSON(http.StatusBadRequest, newErrorResponse("invalid_image", err.Error()))
		}
		buffEffectRepos = append(buffEffectRepos, buffEffectRepo)
	}

	var baseParam float32 = 0
	var boostParam float32 = 0
	for _, buffEffectRepo := range buffEffectRepos {
		baseParam += float32(buffEffectRepo.BuffEffect.BaseParam)
		boostParam += float32(buffEffectRepo.BuffEffect.BoostParam)
	}
	baseParam = (baseParam+15000.0)/250.0 + 1.0
	boostParam = (boostParam + 2650.0) / 10.0
	score := int(baseParam * boostParam)

	skillResponse, err := buildSkillResponse(buffEffectRepos, buffSkillRepo)
	if err != nil {
		log.Printf("Error building skill response: %v", err)
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

func buildSkillResponse(buffEffectRepos []*BuffEffectRepo.BuffEffectRepo, buffSkillRepo *BuffSkillRepo.BuffSkillRepo) ([]SkillResponse, error) {
	var result []SkillResponse
	for _, buffEffectRepo := range buffEffectRepos {
		skill, err := buffSkillRepo.GetSkillFromID(buffEffectRepo.BuffEffect.SkillId)
		if err != nil {
			return nil, err
		}
		result = append(result, SkillResponse{
			Name:       skill.Name,
			BaseParam:  buffEffectRepo.BuffEffect.BaseParam,
			BoostParam: buffEffectRepo.BuffEffect.BoostParam,
		})
	}
	return result, nil
}
