package main

import (
	"log"

	"github.com/itsLeonB/ocram/internal/config"
	"github.com/itsLeonB/ocram/internal/delivery/job"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rotisserie/eris"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(eris.ToString(err, true))
	}

	j, err := job.ExtractExpenseBillTextJob(cfg)
	if err != nil {
		log.Fatal(eris.ToString(err, true))
	}

	j.Run()
}
