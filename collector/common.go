// -*- coding: utf-8 -*-
//
// © Copyright 2023 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace      = "smcipmi"
	CmdSmcIpmiTool = "SMCIPMITool"
)

// Function signature for NewCollector...
type NewCollectorHandle func(string, string, string) prometheus.Collector

type metricTemplate struct {
	desc         *prometheus.Desc
	valueType    prometheus.ValueType
	valueCreator func(string) (float64, error)
}

func createErrorMetric(collector string, target string) prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "collector", "error"),
			"Set if an error has occurred in a collector",
			[]string{"name", "target"},
			nil,
		),
		prometheus.GaugeValue,
		1,
		collector, target,
	)
}
