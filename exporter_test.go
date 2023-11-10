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
		t.Errorf("Incomplete count of pminfo metrics - expected: %d - received: %d", expected, received)
	}

}

func TestPminfoConvertPowerSupplyStatusValue(t *testing.T) {

	testMap := make(map[string]float64) // [PSU-Status]expectedValue
	testMap["                             OK"] = 0.0
	testMap["                Power Supply OK"] = 0.0
	testMap["            [UNIT IS OFF] (40h)"] = 1.0
	testMap["                          (00h)"] = 1.0
	testMap[" [IOUT_OC_FAULT][UNIT IS OFF] (50h)"] = 2.0
	testMap[" [VIN_UV_FAULT][UNIT IS OFF] (48h)"] = 2.0
	testMap["     [Over Current Fault] (08h)"] = 3.0

	for psuStatus, expected := range testMap {
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
