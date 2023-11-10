// -*- coding: utf-8 -*-
//
// © Copyright 2023 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package main

import (
	"prometheus-smcipmi-exporter/collector"
	"prometheus-smcipmi-exporter/util"
	"testing"
)

var (
	testPminfoFile = "pminfo.txt"
)

func TestParsePminfoModule(t *testing.T) {

	var collector collector.PminfoCollector
	pminfoData := util.MustReadFile(&testPminfoFile)
	metrics := collector.CreateMetrics(pminfoData)

	if len(metrics) == 0 {
		t.Error("No pminfo metrics received")
	}

	expected := 14
	received := len(metrics)
	if received != expected {
		t.Errorf("Incomplete count of pminfo metrics - "+
			"expected: %d - received: %d", expected, received)
	}

}

func TestPminfoConvertPowerSupplyStatusValue(t *testing.T) {

	m := make(map[string]float64) // [PSU-Status]expectedValue
	m["                             OK"] = collector.PminfoPsuStateOK
	m["                Power Supply OK"] = collector.PminfoPsuStateOK
	m["            [UNIT IS OFF] (40h)"] = collector.PminfoPsuStateOff
	m["                          (00h)"] = collector.PminfoPsuStateOff
	m["     [Over Current Fault] (08h)"] = collector.PminfoPsuStateFaulty
	m[" [IOUT_OC_FAULT][UNIT IS OFF] (50h)"] = collector.PminfoPsuStateError
	m[" [VIN_UV_FAULT][UNIT IS OFF] (48h)"] = collector.PminfoPsuStateError

	for psuStatus, expected := range m {
		received, err := collector.ConvertPowerSupplyStatusValue(psuStatus)
		if err != nil {
			t.Error("No error expected - but received: ", err)
		}
		if received != float64(expected) {
			t.Errorf("Convertion of PSU value vailed for: %s, "+
				"expected: %v - received: %v", psuStatus, expected, received)
		}
	}
}
