package ImageTextExtractor

import (
	"context"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

// 入力された画像から画像に記載されているテキストを抽出して出力します。
func ExtractTextFromImage(ctx context.Context, file *os.File) (string, error) {
	image, imageErr := vision.NewImageFromReader(file)
	if imageErr != nil {
		return "", imageErr
	}
	client, visionErr := vision.NewImageAnnotatorClient(ctx)
	if visionErr != nil {
		log.Println(visionErr.Error())
		return "", visionErr
	}
	defer client.Close()

	annotations, annotateErr := client.DetectTexts(ctx, image, nil, 500)
	if annotateErr != nil {
		return "", annotateErr
	}

	text := ""
	for _, annotation := range annotations {
		text += annotation.Description
	}

	return text, nil
}
