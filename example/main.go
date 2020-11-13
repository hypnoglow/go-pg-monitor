package main

import (
	"os"

	"github.com/go-pg/pg/v9"

	monitor "github.com/hypnoglow/go-pg-monitor"
	"github.com/hypnoglow/go-pg-monitor/gopgv9"
)

func main() {
	dbOpts, err := pg.ParseURL(os.Getenv("POSTGRES_DSN"))
	if err != nil {
		panic(err)
	}

	db := pg.Connect(dbOpts)

	mon := monitor.NewMonitor(
		gopgv9.NewObserver(db),
		monitor.NewMetrics(monitor.MetricsWithNamespace("my_app")),
	)

	mon.Open()
	defer mon.Close()

	// Set up Prometheus handler for scraping, open db connections, do queries, etc...
}
