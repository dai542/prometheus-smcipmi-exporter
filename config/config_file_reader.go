// Copyright 2023 Gabriele Iannetti <g.iannetti@gsi.de>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

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
