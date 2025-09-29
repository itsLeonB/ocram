package client

import (
	"context"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"github.com/rotisserie/eris"
	"google.golang.org/api/option"
)

type OCRClient interface {
	ExtractFromURI(ctx context.Context, uri string) (string, error)
	Shutdown() error
}

type cloudVisionClient struct {
	client *vision.ImageAnnotatorClient
}

func NewOCRClient(serviceAccountJSON []byte) (OCRClient, error) {
	if serviceAccountJSON == nil {
		return nil, eris.New("service account JSON is nil")
	}

	client, err := vision.NewImageAnnotatorClient(
		context.Background(),
		option.WithCredentialsJSON(serviceAccountJSON),
	)
	if err != nil {
		return nil, eris.Wrap(err, "error initializing vision client")
	}

	return &cloudVisionClient{client}, nil
}

func (cvc *cloudVisionClient) ExtractFromURI(ctx context.Context, uri string) (string, error) {
	img := &visionpb.Image{
		Source: &visionpb.ImageSource{
			GcsImageUri: uri,
		},
	}

	annotation, err := cvc.client.DetectDocumentText(ctx, img, nil)
	if err != nil {
		return "", eris.Wrap(err, "error detecting document text")
	}

	return annotation.GetText(), nil
}

func (cvc *cloudVisionClient) Shutdown() error {
	return cvc.client.Close()
}
