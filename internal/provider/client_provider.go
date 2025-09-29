package provider

import (
	"github.com/itsLeonB/ocram/internal/client"
	"github.com/rotisserie/eris"
)

type Clients struct {
	OCR client.OCRClient
}

func ProvideClients(serviceAccountJSON string) (*Clients, error) {
	if serviceAccountJSON == "" {
		return nil, eris.New("service account JSON cannot be empty")
	}

	ocr, err := client.NewOCRClient([]byte(serviceAccountJSON))
	if err != nil {
		return nil, err
	}

	return &Clients{
		OCR: ocr,
	}, nil
}

func (c *Clients) Shutdown() error {
	return c.OCR.Shutdown()
}
