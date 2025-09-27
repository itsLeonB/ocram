package provider

import (
	"errors"

	"github.com/itsLeonB/ezutil/v2"
	"github.com/itsLeonB/meq"
	"github.com/itsLeonB/ocram/internal/config"
	"github.com/rotisserie/eris"
)

type Providers struct {
	meq.DB
	*Clients
	*Services
}

func ProvideAll(logger ezutil.Logger, cfg config.Config) (*Providers, error) {
	db := meq.NewAsynqDB(logger, cfg.ToRedisOpts())
	queues, err := ProvideQueues(db, logger)
	if err != nil {
		return nil, err
	}
	clients, err := ProvideClients(cfg.ServiceAccount)
	if err != nil {
		return nil, err
	}
	services, err := ProvideServices(clients, queues)
	if err != nil {
		return nil, err
	}

	return &Providers{
		DB:       db,
		Clients:  clients,
		Services: services,
	}, nil
}

func (p *Providers) Ping() error {
	if err := p.DB.Ping(); err != nil {
		return eris.Wrap(err, "error pinging meq.DB")
	}
	return nil
}

func (p *Providers) Shutdown() error {
	var err error
	if e := p.DB.Shutdown(); e != nil {
		err = errors.Join(err, e)
	}
	if e := p.Clients.Shutdown(); e != nil {
		err = errors.Join(err, e)
	}
	return err
}
