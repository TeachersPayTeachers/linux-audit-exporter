# linux-audit-exporter

[linux-audit-exporter](https://github.com/teacherspayteachers/linux-audit-exporter) - Export your Linux Audit status as Prometheus metrics

## TL;DR;

```console
$ helm repo add tpt https://teacherspayteachers.github.io/helm-charts
$ helm repo update
$ helm install linux-audit-exporter tpt/linux-audit-exporter -n kube-system
```

## Introduction

This chart deploys a Prometheus exporter for Linux Audit status on a [Kubernetes](http://kubernetes.io) cluster using the [Helm](https://helm.sh) package manager.

## Prerequisites

- K8S host running Linux kernel with audit subsystem

## Installing the Chart

To install the chart with the release name `linux-audit-exporter`:

```console
$ helm install linux-audit-exporter tpt/linux-audit-exporter -n kube-system
```

The command deploys a Prometheus exporter for Linux Audit status on the Kubernetes cluster in the default configuration. The [configuration](#configuration) section lists the parameters that can be configured during installation.

> **Tip**: List all releases using `helm list`

## Uninstalling the Chart

To uninstall/delete the `linux-audit-exporter`:

```console
$ helm delete linux-audit-exporter -n kube-system
```

The command removes all the Kubernetes components associated with the chart and deletes the release.

## Configuration

The following table lists the configurable parameters of the `linux-audit-exporter` chart and their default values.

|             Parameter              |                                            Description                                            |                  Default                   |
|------------------------------------|---------------------------------------------------------------------------------------------------|--------------------------------------------|
| affinity                           | Specify Pod affinity constraints.                                                                 | `{}`                                       |
| image.pullPolicy                   | Docker image pull policy.                                                                         | `IfNotPresent`                             |
| image.registry                     | Docker image registry.                                                                            | `docker.io`                                |
| image.repository                   | Docker image repository.                                                                          | `teacherspayteachers/linux-audit-exporter` |
| image.tag                          | Docker image tag.                                                                                 | `latest`                                   |
| imagePullSecrets                   | Image pull secrets.                                                                               | `[]`                                       |
| fullnameOverride                   | Full name override.                                                                               | `~`                                        |
| livenessProbe.httpGet.path         | Liveness probe HTTP path.                                                                         | `/healthz`                                 |
| livenessProbe.httpGet.port         | Liveness probe HTTP port.                                                                         | `http`                                     |
| livenessProbe.initialDelaySeconds  | Liveness probe initial delay seconds.                                                             | `10`                                       |
| livenessProbe.periodSeconds        | Liveness probe period seconds.                                                                    | `10`                                       |
| livenessProbe.timeoutSeconds       | Liveness probe timeout seconds.                                                                   | `5`                                        |
| livenessProbe.failureThreshold     | Liveness probe failure threshold.                                                                 | `2`                                        |
| livenessProbe.successThreshold     | Liveness probe success threshold.                                                                 | `1`                                        |
| readinessProbe.httpGet.path        | Readiness probe HTTP path.                                                                        | `/healthz`                                 |
| readinessProbe.httpGet.port        | Readiness probe HTTP port.                                                                        | `http`                                     |
| readinessProbe.initialDelaySeconds | Readiness probe initial delay seconds.                                                            | `5`                                        |
| readinessProbe.periodSeconds       | Readiness probe period seconds.                                                                   | `10`                                       |
| readinessProbe.timeoutSeconds      | Readiness probe timeout seconds.                                                                  | `5`                                        |
| readinessProbe.failureThreshold    | Readiness probe failure threshold.                                                                | `6`                                        |
| readinessProbe.successThreshold    | Readiness probe success threshold.                                                                | `1`                                        |
| nameOverride                       | Name override.                                                                                    | `~`                                        |
| nodeSelector                       | Node selector.                                                                                    | `{}`                                       |
| priorityClassName                  | Priority class name.                                                                              | `""`                                       |
| podAnnotations                     | Pod annotations.                                                                                  | `{}`                                       |
| podLabels                          | Pod labels.                                                                                       | `{}`                                       |
| resources                          | Container resources.                                                                              | `{}`                                       |
| securityContext.readOnlyFileSystem | Set to true to increase attack cost.                                                              | `true`                                     |
| securityContext.runAsUser          | To read from a netlink multicast group, processes require an effective UID of 0 or CAP_NET_ADMIN. | `0`                                        |
| tolerations                        | Specify taint tolerations.                                                                        | `[]`                                       |


Specify each parameter using the `--set key=value[,key=value]` argument to `helm install`. For example:

```console
$ helm install linux-audit-exporter tpt/linux-audit-exporter -n kube-system --set image.pullPolicy=IfNotPresent
```

Alternatively, a YAML file that specifies the values for the parameters can be provided while
installing the chart. For example:

```console
$ helm install linux-audit-exporter tpt/linux-audit-exporter -n kube-system --values values.yaml
```
