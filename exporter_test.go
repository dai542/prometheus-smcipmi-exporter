// -*- coding: utf-8 -*-
//
// © Copyright 2023 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package main

import (
	"flag"
	"os"
	"prometheus-smcipmi-exporter/collector"
	"prometheus-smcipmi-exporter/util"
	"testing"

	log "github.com/sirupsen/logrus"
)

const (
	defaultPminfoFile = "pminfo.txt"
)

var (
	pminfoFile *string
)

// TODO: create local test for exported metrics via HTTP...
func TestParsePminfoModule(t *testing.T) {

	var collector collector.PminfoCollector

	pminfoData := util.MustReadFile(pminfoFile)

	metrics := collector.CreateMetrics(pminfoData)

	for _, m := range metrics {
		log.Debug(m.Desc().String())
	}

	if len(metrics) == 0 {
		t.Error("No pminfo metrics received")
	}

	expected := 14
	received := len(metrics)
	if received != expected {
		t.Errorf("Incomplete count of pminfo metrics - expected: %d - received: %d", expected, received)
	}

}

func TestMain(m *testing.M) {

	pminfoFile = flag.String("pminfoFile", defaultPminfoFile, "A file with pminfo output generated from SMCIPMITool to be processed")
	logLevel := flag.String("log", defaultLogLevel, "Sets log level - ERROR, WARNING, INFO, DEBUG or TRACE")

	flag.Parse()

	initLogging(*logLevel)

	os.Exit(m.Run())
}
