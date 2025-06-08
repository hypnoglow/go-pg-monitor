package main

import (
	"fmt"
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

	// create unique pool name from DB address and database name
	poolName := fmt.Sprintf("%s/%s", dbOpts.Addr, dbOpts.Database)

	mon := monitor.NewMonitor(
		// Observer package must match your go-pg version.
		// E.g. for go-pg v10.x.x use package gopgv10.
		gopgv9.NewObserver(db),
		monitor.NewMetrics(
			// If you already have application labels on your scraped metrics (e.g. in k8s),
			// then you probably don't need this, because you can query metrics with selector:
			// go_pg_pool_hits{app="my-app"}
			//
			// If your metrics lack such labels then you can set this option,
			// so your metrics will look like this:
			// my_app_go_pg_pool_hits{}
			monitor.MetricsWithNamespace("my_app"),
		),
		// set pool name
		monitor.MonitorWithPoolName(poolName),
	)

	mon.Open()
	defer mon.Close()

	// Set up Prometheus handler for scraping, open db connections, do queries, etc...
}
