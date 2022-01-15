package exporter

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"

	"github.com/TeachersPayTeachers/linux-audit-exporter/internal/audit"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	c := audit.NewMockClient(ctrl)

	_ = New(c)
}

func TestDescribe(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	c := audit.NewMockClient(ctrl)

	e := New(c)

	ch := make(chan *prometheus.Desc, 100)

	e.Describe(ch)

	close(ch)

	desc := make([]*prometheus.Desc, 0)

	for d := range ch {
		desc = append(desc, d)
	}

	if len(desc) != 8 {
		t.Errorf("Got %d metric descriptions; want 8.", len(desc))
	}

	// Backlog.

	expect := fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_backlog",
		"Number of event records currently queued waiting for auditd to read them.",
	)
	if desc[0].String() != expect {
		t.Errorf("Got %s metric[0] description; want %s.", desc[0].String(), expect)
	}

	// Backlog limit.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_backlog_limit",
		"Number of outstanding audit buffers allowed.",
	)
	if desc[1].String() != expect {
		t.Errorf("Got %s metric[1] description; want %s.", desc[1].String(), expect)
	}

	// Backlog wait time.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_backlog_wait_time",
		"Time kernel waits when backlog limit is reached.",
	)
	if desc[2].String() != expect {
		t.Errorf("Got %s metric[2] description; want %s.", desc[2].String(), expect)
	}

	// Backlog wait time actual.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_backlog_wait_time_actual",
		"Total time spent by kernel waiting to queue audit events on backlog.",
	)
	if desc[3].String() != expect {
		t.Errorf("Got %s metric[3] description; want %s.", desc[3].String(), expect)
	}

	// Enabled.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_enabled",
		"Enabled flag. 0 = disabled. 1 = enabled. 2 = immutable. -1 = unknown.",
	)
	if desc[4].String() != expect {
		t.Errorf("Got %s metric[3] description; want %s.", desc[4].String(), expect)
	}

	// Backlog limit.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_failure",
		"Number of critical errors, such as transmission errors, backlog limit exceeded, etc.",
	)
	if desc[5].String() != expect {
		t.Errorf("Got %s metric[4] description; want %s.", desc[5].String(), expect)
	}

	// Lost.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_lost",
		"Number of event records that have been discarded due to kernel audit queue overflowing.",
	)
	if desc[6].String() != expect {
		t.Errorf("Got %s metric[6] description; want %s.", desc[6].String(), expect)
	}

	// Rate limit.

	expect = fmt.Sprintf(
		"Desc{fqName: \"%s\", help: \"%s\", constLabels: {}, variableLabels: []}",
		"linux_audit_rate_limit",
		"Limit of messages per second. A value of zero means no rate limit is applied.",
	)
	if desc[7].String() != expect {
		t.Errorf("Got %s metric[6] description; want %s.", desc[7].String(), expect)
	}
}

func TestCollect(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	status := &audit.Status{
		Backlog:               100,
		BacklogLimit:          5000,
		BacklogWaitTime:       30000,
		BacklogWaitTimeActual: 2500,
		Enabled:               1,
		Failure:               150,
		Lost:                  15,
		RateLimit:             1500,
	}

	c := audit.NewMockClient(ctrl)

	c.EXPECT().GetStatus().Return(status, nil)

	e := New(c)

	ch := make(chan prometheus.Metric, 100)

	e.Collect(ch)

	close(ch)

	metric := make([]prometheus.Metric, 0)

	for m := range ch {
		metric = append(metric, m)
	}

	if len(metric) != 8 {
		t.Errorf("Got %d metrics; want 8.", len(metric))
	}

	// Backlog.

	// nolint:exhaustivestruct
	o := &dto.Metric{}

	if metric[0].Desc() != e.backlogDesc {
		t.Errorf("Got metric[0].Desc() %s; want %s.", metric[0].Desc().String(), e.backlogDesc.String())
	}

	if err := metric[0].Write(o); err != nil {
		t.Errorf("Failed to write metric[0].")
	}

	if o.Gauge.GetValue() != float64(status.Backlog) {
		t.Errorf("Got metric[0] value %f; wanted %d.", o.Gauge.GetValue(), status.Backlog)
	}

	// Backlog limit.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[1].Desc() != e.backlogLimitDesc {
		t.Errorf("Got metric[1].Desc() %s; want %s.", metric[1].Desc().String(), e.backlogLimitDesc.String())
	}

	if err := metric[1].Write(o); err != nil {
		t.Errorf("Failed to write metric[1].")
	}

	if o.Gauge.GetValue() != float64(status.BacklogLimit) {
		t.Errorf("Got metric[1] value %f; wanted %d.", o.Gauge.GetValue(), status.BacklogLimit)
	}

	// Backlog wait time.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[2].Desc() != e.backlogWaitTimeDesc {
		t.Errorf("Got metric[2].Desc() %s; want %s.", metric[2].Desc().String(), e.backlogWaitTimeDesc.String())
	}

	if err := metric[2].Write(o); err != nil {
		t.Errorf("Failed to write metric[2].")
	}

	if o.Gauge.GetValue() != float64(status.BacklogWaitTime) {
		t.Errorf("Got metric[2] value %f; wanted %d.", o.Gauge.GetValue(), status.BacklogWaitTime)
	}

	// Backlog wait time actual.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[3].Desc() != e.backlogWaitTimeActualDesc {
		t.Errorf("Got metric[3].Desc() %s; want %s.", metric[3].Desc().String(), e.backlogWaitTimeActualDesc.String())
	}

	if err := metric[3].Write(o); err != nil {
		t.Errorf("Failed to write metric[3].")
	}

	if o.Gauge.GetValue() != float64(status.BacklogWaitTimeActual) {
		t.Errorf("Got metric[3] value %f; wanted %d.", o.Gauge.GetValue(), status.BacklogWaitTimeActual)
	}

	// Enabled.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[4].Desc() != e.enabledDesc {
		t.Errorf("Got metric[4].Desc() %s; want %s.", metric[4].Desc().String(), e.enabledDesc.String())
	}

	if err := metric[4].Write(o); err != nil {
		t.Errorf("Failed to write metric[4].")
	}

	if o.Gauge.GetValue() != float64(status.Enabled) {
		t.Errorf("Got metric[4] value %f; wanted %d.", o.Gauge.GetValue(), status.Enabled)
	}

	// Failure.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[5].Desc() != e.failureDesc {
		t.Errorf("Got metric[5].Desc() %s; want %s.", metric[5].Desc().String(), e.failureDesc.String())
	}

	if err := metric[5].Write(o); err != nil {
		t.Errorf("Failed to write metric[5].")
	}

	if o.Gauge.GetValue() != float64(status.Failure) {
		t.Errorf("Got metric[5] value %f; wanted %d.", o.Gauge.GetValue(), status.Failure)
	}

	// Lost.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[6].Desc() != e.lostDesc {
		t.Errorf("Got metric[6].Desc() %s; want %s.", metric[6].Desc().String(), e.lostDesc.String())
	}

	if err := metric[6].Write(o); err != nil {
		t.Errorf("Failed to write metric[6].")
	}

	if o.Gauge.GetValue() != float64(status.Lost) {
		t.Errorf("Got metric[6] value %f; wanted %d.", o.Gauge.GetValue(), status.Lost)
	}

	// Rate limit.

	// nolint:exhaustivestruct
	o = &dto.Metric{}

	if metric[7].Desc() != e.rateLimitDesc {
		t.Errorf("Got metric[7].Desc() %s; want %s.", metric[7].Desc().String(), e.rateLimitDesc.String())
	}

	if err := metric[7].Write(o); err != nil {
		t.Errorf("Failed to write metric[7].")
	}

	if o.Gauge.GetValue() != float64(status.RateLimit) {
		t.Errorf("Got metric[7] value %f; wanted %d.", o.Gauge.GetValue(), status.RateLimit)
	}
}
