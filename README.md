linux-audit-exporter
====================

A Prometheus exporter for Linux Audit status.

## Project Status

This project is:

 * In ALPHA state
 * Is maintained by Teachers Pay Teachers
 * Is used in production by Teachers Pay Teachers

### Versioning

While this project is in ALPHA status, breaking changes
may occur between minor versions, and will be announced
in [CHANGELOG.md](https://github.com/TeachersPayTeachers/linux-audit-exporter/blob/main/CHANGELOG.md)
and in the GitHub release notes.

The API surface area includes:

 * Runtime dependencies
 * The Helm chart values
 * The CLI flags and environment variables
 * The exported Prometheus metric names and labels

## Requirements

Golang 1.14+ is required to build the project.

Will run on any system that can run Go binaries, but will only
export Linux Audit status when run from a Linux host. When run
from a non-Linux host, it will export zero-ed metrics.

Additionally, requires:

 * Root (UID 0) user to run binary
 * Audit read privilege (`CAP_AUDIT_READ`)
 * `--host=pid` in Docker, `hostPID: true` in Kubernetes

## Why

The Linux audit system can be configured to log security-relevant
events such as syscalls, and processed by auditd or commercial
alternatives.

On high-throughput systems it may generate audit records at a faster
pace than auditd (or equivalent) can consume. Depending on how
Linux audit is configured, this can result in:

 * Lost audit events
 * Backlog processing delays
 * Out-of-memory

The linux-audit-exporter surfaces metrics meant to help operators
monitor the status of the Linux audit system, and avoid these
outcomes.

## Build

### Go binary

```
$ go build .
```

### Docker image

```
$ docker build .
```

## Usage

### Help

```
$ linux-audit-exporter -h
Export Linux audit status as Prometheus metrics

Usage:
  linux-audit-exporter [flags]

Flags:
      --health-path string      health path (default "/healthz")
  -h, --help                    help for linux-audit-exporter
      --listen-address string   listen address (default "0.0.0.0:9090")
      --metrics-path string     metrics path (default "/metrics")
```

### Run

Run binary directly:

```
$ linux-audit-exporter
```

Or, use Docker image:

```
$ docker run -p 9090:9090 --privileged TeachersPayTeachers/linux-audit-exporter
```

### Get metrics

```
$ curl localhost:9090/metrics | grep linux_audit
```

## Deploy

### Docker

Docker images are published [here](https://hub.docker.com/repository/docker/teacherspayteachers/linux-audit-exporter).

```
$ docker run --privileged teacherspayteachers/linux-audit-exporter:latest
```

### Helm

A Helm chart is available, which deploys linux-audit-exporter as a DaemonSet.

```
$ helm repo add tpt https://teacherspayteachers.github.io/helm-charts
$ helm install linux-audit-exporter tpt/linux-audit-exporter
```

The Helm chart code is located [here](https://github.com/TeachersPayTeachers/linux-audit-exporter/tree/main/deploy/helm-charts/linux-audit-exporter).

## Metrics

When run from a Linux host with required privileges, will export
Linux Audit status as Prometheus metrics. Here is a sample:

```
# HELP linux_audit_backlog Number of event records currently queued waiting for auditd to read them.
# TYPE linux_audit_backlog gauge
linux_audit_backlog 0
# HELP linux_audit_backlog_limit Number of outstanding audit buffers allowed.
# TYPE linux_audit_backlog_limit gauge
linux_audit_backlog_limit 5000
# HELP linux_audit_backlog_wait_time Time kernel waits when backlog limit is reached.
# TYPE linux_audit_backlog_wait_time gauge
linux_audit_backlog_wait_time 15000
# HELP linux_audit_backlog_wait_time_actual Total time spent by kernel waiting to queue audit events on backlog.
# TYPE linux_audit_backlog_wait_time_actual gauge
linux_audit_backlog_wait_time_actual 10
# HELP linux_audit_enabled Enabled flag. 0 = disabled. 1 = enabled. 2 = immutable. -1 = unknown.
# TYPE linux_audit_enabled gauge
linux_audit_enabled 1
# HELP linux_audit_failure Number of critical errors, such as transmission errors, backlog limit exceeded, etc.
# TYPE linux_audit_failure gauge
linux_audit_failure 0
# HELP linux_audit_lost Number of event records that have been discarded due to kernel audit queue overflowing.
# TYPE linux_audit_lost gauge
linux_audit_lost 0
# HELP linux_audit_rate_limit Limit of messages per second. A value of zero means no rate limit is applied.
# TYPE linux_audit_rate_limit gauge
linux_audit_rate_limit 100000
```

## Alternatives

### `printk` and `dmesg`

The Linux audit system can be configured to log failures and lost
events with `printk`, which can usually be read with `dmesg`.

### StatsD plugin

There is an (experimental, at time of this writing) user-space [StatsD plugin](https://github.com/linux-audit/audit-userspace/tree/27bb97d1dd04cc3768ab7756008a7164f308bf85/audisp/plugins/statsd).

## Contributing

Contributions are very welcome! Please see [CONTRIBUTING.md](https://github.com/TeachersPayTeachers/linux-audit-exporter/blob/main/CONTRIBUTING.md).

## License

[MIT](https://github.com/TeachersPayTeachers/linux-audit-exporter/blob/main/LICENSE.md)
