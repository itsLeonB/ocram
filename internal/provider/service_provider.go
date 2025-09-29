package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/ocram/internal/service"
	"github.com/rotisserie/eris"
)

type Services struct {
	ExpenseBill service.ExpenseBillService
}

func ProvideServices(clients *Clients, queues *Queues, logger ezutil.Logger) (*Services, error) {
	if clients == nil {
		return nil, eris.New("clients cannot be nil")
	}
	if queues == nil {
		return nil, eris.New("queues cannot be nil")
	}
	if logger == nil {
		return nil, eris.New("logger cannot be nil")
	}

	expenseBill, err := service.NewExpenseBillService(
		queues.ExpenseBillTextExtracted,
		queues.ExpenseBillUploaded,
		clients.OCR,
		logger,
	)
	if err != nil {
		return nil, err
	}

	return &Services{ExpenseBill: expenseBill}, nil
}
