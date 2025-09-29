package provider

import (
	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/ocram/internal/message"
	"github.com/rotisserie/eris"
)

type Queues struct {
	ExpenseBillTextExtracted meq.TaskQueue[message.ExpenseBillTextExtracted]
	ExpenseBillUploaded      meq.TaskQueue[message.ExpenseBillUploaded]
}

func ProvideQueues(db meq.DB, logger ezutil.Logger) (*Queues, error) {
	if db == nil {
		return nil, eris.New("db cannot be nil")
	}
	if logger == nil {
		return nil, eris.New("logger cannot be nil")
	}
	return &Queues{
		ExpenseBillTextExtracted: meq.NewTaskQueue[message.ExpenseBillTextExtracted](logger, db),
		ExpenseBillUploaded:      meq.NewTaskQueue[message.ExpenseBillUploaded](logger, db),
	}, nil
}
