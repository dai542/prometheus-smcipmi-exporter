# Prometheus SMCIPMI Exporter

A Prometheus exporter to expose metrics specific to the [Supermicro IPMI Utilities SMCIPMITool](https://www.supermicro.com/de/solutions/management-software/ipmi-utilities).

Why an exporter for SMCIPMITool?  

Because FreeIPMI cannot access power supply information of a JBOD system provided by Supermicro.

## Build

`go build`

## Requirements

[Supermicro IPMI Utilities](https://www.supermicro.com/en/solutions/management-software/ipmi-utilities)

## Metrics

All metrics are prefixed with "smcipmi_".

### collector

General metrics about a collector e.g. for pminfo module.

Metrics have an additional prefix "collector_".

| Metric | Labels       | Description                                 |
| -------|------------- | ------------------------------------------- |
| error  | name, target | Set if an error has occurred in a collector |

### pminfo

Metrics specific to the pminfo module.

Currently only a restricted set of metrics for the pminfo module are exposed.

Metrics have an additional prefix "pminfo_".

| Metric                     | Labels         | Description                                  |
| -------------------------- | -------------- | -------------------------------------------- |
| power\_consumption\_watts  | target, module | Current power consumption measured in watts  |
| power\_supply\_status      | target, module | Power supply status (0=OK, 1=OFF, 2=Failure) |
