package sql_metrics

import (
	"database/sql"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	labelQuery = "query"
)

type DBMetrics struct {
	db *sql.DB

	duration      *prometheus.SummaryVec
	errorsTotal   *prometheus.CounterVec
	requestsTotal *prometheus.CounterVec
}

func InitDBMetrics(db *sql.DB) *DBMetrics {
	m := &DBMetrics{
		db: db,
	}

	m.duration = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name:       "duration_seconds",
			Help:       "Summary of sql query duration in seconds",
			Objectives: map[float64]float64{0.9: 0.01, 0.95: 0.005, 0.99: 0.001},
		},
		[]string{labelQuery},
	)

	m.errorsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "errors_total",
			Help: "Total number of errors",
		},
		[]string{labelQuery},
	)

	m.requestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "request_total",
			Help: "Total number of requests",
		},
		[]string{labelQuery},
	)

	prometheus.MustRegister(m.duration)
	prometheus.MustRegister(m.errorsTotal)
	prometheus.MustRegister(m.requestsTotal)

	return m
}

func (m *DBMetrics) Exec(query string, args ...interface{}) (sql.Result, error) {
	t := time.Now()
	res, err := m.db.Exec(query, args)
	if err != nil {
		m.errorsTotal.With(prometheus.Labels{labelQuery: query}).Inc()
	}
	m.requestsTotal.With(prometheus.Labels{labelQuery: query}).Inc()
	m.duration.With(prometheus.Labels{labelQuery: query}).Observe(time.Since(t).Seconds())

	return res, err
}

func (m *DBMetrics) Query(query string, args ...interface{}) (*sql.Rows, error) {
	t := time.Now()
	res, err := m.db.Query(query, args)
	if err != nil {
		m.errorsTotal.With(prometheus.Labels{labelQuery: query}).Inc()
	}
	m.requestsTotal.With(prometheus.Labels{labelQuery: query}).Inc()
	m.duration.With(prometheus.Labels{labelQuery: query}).Observe(time.Since(t).Seconds())

	return res, err
}

func (m *DBMetrics) QueryRow(query string, args ...interface{}) *sql.Row {
	t := time.Now()
	res := m.db.QueryRow(query, args)

	m.requestsTotal.With(prometheus.Labels{labelQuery: query}).Inc()
	m.duration.With(prometheus.Labels{labelQuery: query}).Observe(time.Since(t).Seconds())

	return res
}
