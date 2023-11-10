// -*- coding: utf-8 -*-
//
// © Copyright 2023 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package config

import (
	"log"
	"strconv"
)

type ConfigFileReader struct {
	Targets  []string
	User     string
	Password string

	CollectPminfo bool

	GenericConfigFileReader
}

func NewConfigFileReader(filepath string) *ConfigFileReader {
	c := new(ConfigFileReader)

	c.MustLoadFile(filepath)

	// TODO Check values...
	c.Targets = c.MustHaveStringList("targets")
	c.User = c.MustHaveString("login.user")
	c.Password = c.MustHaveString("login.password")

	collectorMap := c.MustHaveMap("collectors")

	c.CollectPminfo = c.mustHaveCollectorOption(collectorMap, "pminfo")

	return c
}

func (ConfigFileReader) mustHaveCollectorOption(m map[string]string, key string) bool {

	value, ok := m[key]
	if !ok {
		log.Panicf("Collector option %s not found", key)
	}
	if value == "" {
		log.Panicf("Collector option %s has no value", key)
	}
	collect, err := strconv.ParseBool(value)
	if err != nil {
		log.Panicf("Error converting value for %s collector option...\n", err)
	}

	return collect
}
