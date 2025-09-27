package job

import (
	"context"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/ocram/internal/config"
	"github.com/itsLeonB/ocram/internal/provider"
	"github.com/itsLeonB/ocram/internal/service"
)

type extractExpenseBillTextJob struct {
	expenseBillSvc service.ExpenseBillService
}

func (j *extractExpenseBillTextJob) Run() error {
	return j.expenseBillSvc.ExtractBillText(context.Background())
}

func ExtractExpenseBillTextJob(cfg config.Config) (*ezutil.Job, error) {
	logger := provider.ProvideLogger(config.AppName, cfg.Env)
	providers, err := provider.ProvideAll(logger, cfg)
	if err != nil {
		return nil, err
	}

	jobImpl := extractExpenseBillTextJob{providers.ExpenseBill}

	return ezutil.NewJob(logger, jobImpl.Run).
		WithSetupFunc(providers.Ping).
		WithCleanupFunc(providers.Shutdown), nil
}
