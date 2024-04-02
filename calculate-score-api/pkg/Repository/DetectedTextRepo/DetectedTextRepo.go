package DetectedTextRepo

import (
	"context"
	"errors"
	"image"
	"log"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
	"manntera.com/calculate-score-api/pkg/NormalizeRect"
	"manntera.com/calculate-score-api/pkg/Repository/SamplerImageRepo"
)

type DetectedTextRepo struct {
	DetectedText []*DetectedText `json:"detectedText"`
}

var _ IDetectedTextRepo = &DetectedTextRepo{}

func NewDetectedTextRepoFromSamplerImageRepo(ctx context.Context, samplerImageRepo *SamplerImageRepo.SamplerImageRepo) (*DetectedTextRepo, error) {
	visionImage, imageErr := vision.NewImageFromReader(samplerImageRepo.GetFile())
	if imageErr != nil {
		return nil, imageErr
	}
	client, visionErr := vision.NewImageAnnotatorClient(ctx)
	if visionErr != nil {
		return nil, visionErr
	}
	defer client.Close()

	annotations, annotateErr := client.DetectTexts(ctx, visionImage, nil, 500)
	if annotateErr != nil {
		return nil, annotateErr
	}

	if len(annotations) == 0 {
		return nil, errors.New("annotations are empty")
	}
	result := &DetectedTextRepo{}
	detectedTexts := make([]*DetectedText, 0)
	lines := strings.Split(annotations[0].Description, "\n")

	srcImageRect := samplerImageRepo.GetImageSize()
	annotationIndex := 1
	for _, line := range lines {
		line = strings.ReplaceAll(line, " ", "")
		line = strings.ReplaceAll(line, "　", "")
		log.Printf("【LINE】 %s", line)
		if annotationIndex >= len(annotations) {
			break
		}
		annotation := annotations[annotationIndex]
		minPoint := image.Point{
			X: int(annotation.BoundingPoly.Vertices[0].X),
			Y: int(annotation.BoundingPoly.Vertices[0].Y),
		}
		maxPoint := image.Point{
			X: int(annotation.BoundingPoly.Vertices[2].X),
			Y: int(annotation.BoundingPoly.Vertices[2].Y),
		}
		detectedText := DetectedText{
			Text: "",
			Rect: image.Rectangle{
				Min: minPoint,
				Max: maxPoint,
			},
			NormalizeRect: NormalizeRect.NormalizeRect{
				Min: NormalizeRect.NormalizedPoint{
					X: float64(minPoint.X) / float64(srcImageRect.Width),
					Y: float64(minPoint.Y) / float64(srcImageRect.Height),
				},
				Max: NormalizeRect.NormalizedPoint{
					X: float64(maxPoint.X) / float64(srcImageRect.Width),
					Y: float64(maxPoint.Y) / float64(srcImageRect.Height),
				},
			},
		}
		for annotationIndex < len(annotations) {
			if detectedText.Text == line {
				break
			}
			if annotationIndex >= len(annotations) {
				break
			}
			log.Printf("  -【ANNO】 %s", annotations[annotationIndex].Description)
			annotation := annotations[annotationIndex]
			detectedText.Text += annotation.Description
			detectedText.Rect.Max.X = int(annotation.BoundingPoly.Vertices[2].X)
			detectedText.Rect.Max.Y = int(annotation.BoundingPoly.Vertices[2].Y)
			detectedText.NormalizeRect.Max.X = float64(detectedText.Rect.Max.X) / float64(srcImageRect.Width)
			detectedText.NormalizeRect.Max.Y = float64(detectedText.Rect.Max.Y) / float64(srcImageRect.Height)
			annotationIndex++
		}
		detectedTexts = append(detectedTexts, &detectedText)
	}
	result.DetectedText = detectedTexts
	return result, nil
}

func (repo *DetectedTextRepo) FindLineTextFromKeyword(keyword string, searchNormalizeRect NormalizeRect.NormalizeRect) ([]*DetectedText, error) {
	result := make([]*DetectedText, 0)
	for _, detectedText := range repo.DetectedText {
		if NormalizeRect.IsCollisionRect(detectedText.NormalizeRect, searchNormalizeRect) {
			if strings.Contains(detectedText.Text, keyword) {
				result = append(result, detectedText)
			}
		}
	}
	return result, nil
}
