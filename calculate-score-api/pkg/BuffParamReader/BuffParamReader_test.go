package BuffParamReader

import (
	"context"
	"os"
	"testing"

	Imagetextextractor "manntera.com/calculate-score-api/pkg/ImageTextExtractor"
)

func TestGetBuffParam(t *testing.T) {
	ctx := context.Background()
	dir, err := os.Getwd()
	if err != nil {
		t.Errorf("Error getting current directory: %v", err)
		return
	}
	testImagePath := dir + "/../../testdata/test4.jpg"

	file, err := os.Open(testImagePath)
	if err != nil {
		t.Errorf("Error opening test image: %v", err)
		return
	}
	defer file.Close()

	text, err := Imagetextextractor.ExtractTextFromImage(ctx, file)
	t.Log(text)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	// buffParam, buffParamErr := GetBuffParams(text, "ラブリーテンポ")
	// buffParam, buffParamErr := GetBuffParams(text, "有名税")
	// buffParam, buffParamErr := GetBuffParams(text, "特別な物語")
	buffParam, buffParamErr := GetBuffParams(text, "アドレナリン")
	if buffParamErr != nil {
		t.Errorf("Error: %v", buffParamErr)
	}
	if len(buffParam) == 0 {
		t.Errorf("Expected buffParam to be non-empty")
	}
	t.Log(buffParam)
}
