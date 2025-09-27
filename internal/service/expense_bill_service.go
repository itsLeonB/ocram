package service

import (
	"context"

	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/ocram/internal/client"
	"github.com/itsLeonB/ocram/internal/config"
	"github.com/itsLeonB/ocram/internal/message"
	"github.com/rotisserie/eris"
)

type ExpenseBillService interface {
	ExtractBillText(ctx context.Context) error
}

type expenseBillServiceImpl struct {
	textExtractedQueue meq.TaskQueue[message.ExpenseBillTextExtracted]
	uploadedQueue      meq.TaskQueue[message.ExpenseBillUploaded]
	ocr                client.OCRClient
}

func NewExpenseBillService(
	textExtractedQueue meq.TaskQueue[message.ExpenseBillTextExtracted],
	uploadedQueue meq.TaskQueue[message.ExpenseBillUploaded],
	ocr client.OCRClient,
) (ExpenseBillService, error) {
	if textExtractedQueue == nil || uploadedQueue == nil {
		return nil, eris.New("queue cannot be nil")
	}
	if ocr == nil {
		return nil, eris.New("ocr client cannot be nil")
	}
	return &expenseBillServiceImpl{
		textExtractedQueue,
		uploadedQueue,
		ocr,
	}, nil
}

func (ebs *expenseBillServiceImpl) ExtractBillText(ctx context.Context) error {
	task, taskID, err := ebs.uploadedQueue.GetOldest(ctx)
	if err != nil {
		return err
	}

	text, err := ebs.ocr.ExtractFromURI(ctx, task.Message.URI)
	if err != nil {
		return err
	}

	msg := message.ExpenseBillTextExtracted{
		ID:   task.Message.ID,
		Text: text,
	}

	if err := ebs.textExtractedQueue.Enqueue(ctx, config.AppName, msg); err != nil {
		return err
	}

	return ebs.uploadedQueue.Delete(ctx, taskID)
}
