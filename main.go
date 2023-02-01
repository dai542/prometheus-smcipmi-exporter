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

package main

import (
	"flag"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

var (
	configFile       *string
	configFileReader *ConfigFileReader
	collectorCreator map[string]NewCollectorHandle
)

const (
	defaultConfigFile = "config.yml"
	defaultPort       = "9850"
)

type NewCollectorHandle func(string, Login) prometheus.Collector

func main() {
	configFile = flag.String("configFile", defaultConfigFile, "Path to YAML config file")
	port := flag.String("port", defaultPort, "The port to listen on for HTTP requests")
	logLevel := flag.String("log", defaultLogLevel, "Sets log level - ERROR, WARNING, INFO, DEBUG or TRACE")

	flag.Parse()

	initLogging(*logLevel)

	configFileReader = newConfigFileReader(*configFile)

	collectorCreator = make(map[string]NewCollectorHandle)

	if configFileReader.collectPminfo {
		collectorCreator["pminfo"] = newPminfoCollector
	}

	targetCount := len(configFileReader.targets)

	for name, collector := range collectorCreator {

		log.Debug("Enable collector: ", name)

		for i := 0; i < targetCount; i++ {
			prometheus.MustRegister(collector(configFileReader.targets[i], configFileReader.login))
		}
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*port, nil)
}
