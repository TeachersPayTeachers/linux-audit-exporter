// Package exporter has logic to export Prometheus metrics.
package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/TeachersPayTeachers/linux-audit-exporter/internal/audit"
)

// Namespace for all metrics.
const prometheusNamespace = "linux_audit"

// Exporter implements prometheus.Collector.
type Exporter struct {
	auditClient         audit.Client
	backlogDesc         *prometheus.Desc
	backlogLimitDesc    *prometheus.Desc
	backlogWaitTimeDesc *prometheus.Desc
	enabledDesc         *prometheus.Desc
	failureDesc         *prometheus.Desc
	lostDesc            *prometheus.Desc
	rateLimitDesc       *prometheus.Desc
}

func New(auditClient audit.Client) *Exporter {
	return &Exporter{
		auditClient: auditClient,
		backlogDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "backlog"),
			"Number of event records currently queued waiting for auditd to read them.",
			[]string{},
			nil,
		),
		backlogLimitDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "backlog_limit"),
			"Number of outstanding audit buffers allowed.",
			[]string{},
			nil,
		),
		backlogWaitTimeDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "backlog_wait_time"),
			"Time kernel waits when backlog limit is reached.",
			[]string{},
			nil,
		),
		enabledDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "enabled"),
			"Enabled flag. 0 = disabled. 1 = enabled. 2 = immutable. -1 = unknown.",
			[]string{},
			nil,
		),
		failureDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "failure"),
			"Number of critical errors, such as transmission errors, backlog limit exceeded, etc.",
			[]string{},
			nil,
		),
		lostDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "lost"),
			"Number of event records that have been discarded due to kernel audit queue overflowing.",
			[]string{},
			nil,
		),
		rateLimitDesc: prometheus.NewDesc(
			prometheus.BuildFQName(prometheusNamespace, "", "rate_limit"),
			"Limit of messages per second. A value of zero means no rate limit is applied.",
			[]string{},
			nil,
		),
	}
}

// Describe satisfies prometheus.Collector interface by sending descriptions
// for all metrics the exporter can possibly report.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	log.Printf("Describing metrics.")
	ch <- e.backlogDesc
	ch <- e.backlogLimitDesc
	ch <- e.backlogWaitTimeDesc
	ch <- e.enabledDesc
	ch <- e.failureDesc
	ch <- e.lostDesc
	ch <- e.rateLimitDesc
}

// Collect satisfies prometheus.Collector interface and sends all metrics.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Printf("Collecting metrics.")

	s, err := e.auditClient.GetStatus()
	if err == nil {
		ch <- prometheus.MustNewConstMetric(e.backlogDesc, prometheus.GaugeValue, float64(s.Backlog))
		ch <- prometheus.MustNewConstMetric(e.backlogLimitDesc, prometheus.GaugeValue, float64(s.BacklogLimit))
		ch <- prometheus.MustNewConstMetric(e.backlogWaitTimeDesc, prometheus.GaugeValue, float64(s.BacklogWaitTime))
		ch <- prometheus.MustNewConstMetric(e.enabledDesc, prometheus.GaugeValue, float64(s.Enabled))
		ch <- prometheus.MustNewConstMetric(e.failureDesc, prometheus.CounterValue, float64(s.Failure))
		ch <- prometheus.MustNewConstMetric(e.lostDesc, prometheus.CounterValue, float64(s.Lost))
		ch <- prometheus.MustNewConstMetric(e.rateLimitDesc, prometheus.GaugeValue, float64(s.RateLimit))
	} else {
		log.Printf("Error getting audit status: %s", err)
	}
}
