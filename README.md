# go-pg-monitor

Provides Prometheus metrics for [go-pg](https://github.com/go-pg/pg) database client.

Example Grafana dashboards:

![grafana](.github/images/grafana.png)

## Usage

0. Your application should already be using go-pg and exporting Prometheus metrics.

1. Add the module to your project:

    ```shell
    go get github.com/hypnoglow/go-pg-monitor
    ```

2. Add monitor to your application entrypoint:

    ```go
    // Create monitor based on your *pg.DB
    mon := monitor.NewMonitor(
		gopgv10.NewObserver(db),
		monitor.NewMetrics(),
	)
   
    // Call this on application startup. 
    mon.Open()
	
    // Call this on application shutdown.
    mon.Close()
    ```

See [example](example/main.go) for details.
