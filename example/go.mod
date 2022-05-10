module github.com/hypnoglow/go-pg-monitor/example

go 1.13

replace (
	github.com/hypnoglow/go-pg-monitor => ../
	github.com/hypnoglow/go-pg-monitor/gopgv9 => ../gopgv9
)

require (
	github.com/go-pg/pg/v9 v9.0.0
	github.com/hypnoglow/go-pg-monitor v1.0.0
	github.com/hypnoglow/go-pg-monitor/gopgv9 v1.0.0
)
