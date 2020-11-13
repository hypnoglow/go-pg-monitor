package monitor

import (
	"context"
	"sync"
	"time"
)

const (
	// defaultInterval is default interval for observing metrics.
	defaultInterval = time.Second * 10
)

// Monitor is a go-pg database client monitor.
type Monitor struct {
	observer Observer
	metrics  *Metrics

	interval time.Duration

	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

// MonitorOption is an option for NewMonitor.
type MonitorOption func(monitor *Monitor)

// NewMonitor returns a new configured Monitor.
func NewMonitor(observer Observer, metrics *Metrics, opts ...MonitorOption) *Monitor {
	m := &Monitor{
		observer: observer,
		metrics:  metrics,
		interval: defaultInterval,
	}

	for _, opt := range opts {
		opt(m)
	}

	return m
}

// Open starts the monitor, observing metrics at each interval.
func (m *Monitor) Open() {
	m.ctx, m.cancel = context.WithCancel(context.Background())

	m.wg.Add(1)
	go func() {
		defer m.wg.Done()

		m.worker()
	}()
}

// Close stops the monitor.
func (m *Monitor) Close() {
	m.cancel()
	m.wg.Wait()
}

func (m *Monitor) worker() {
	// Do first sync right after start.
	m.sync()

	ticker := time.NewTicker(m.interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			m.sync()
		case <-m.ctx.Done():
			return
		}
	}
}

func (m *Monitor) sync() {
	stats := m.observer.Observe()

	m.metrics.hits.Set(float64(stats.Hits))
	m.metrics.misses.Set(float64(stats.Misses))
	m.metrics.timeouts.Set(float64(stats.Timeouts))

	m.metrics.totalConns.Set(float64(stats.TotalConns))
	m.metrics.idleConns.Set(float64(stats.IdleConns))
	m.metrics.staleConns.Set(float64(stats.StaleConns))
}
