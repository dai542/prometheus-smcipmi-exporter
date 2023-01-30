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
	configFileReader ConfigFileReader
)

const (
	defaultConfigFile = "config.yml"
	defaultPort       = "9850"
)

func register(target string, login *Login) {
	// TODO: Modulize collectors
	c := newPminfoCollector(target, login)
	prometheus.MustRegister(c)
}

func main() {
	configFile = flag.String("configFile", defaultConfigFile, "Path to YAML config file")
	port := flag.String("port", defaultPort, "The port to listen on for HTTP requests")
	logLevel := flag.String("log", defaultLogLevel, "Sets log level - ERROR, WARNING, INFO, DEBUG or TRACE")

	flag.Parse()

	initLogging(*logLevel)

	configFileReader.LoadFile(*configFile)

	login := newLogin(
		configFileReader.MustHaveString("login.user"),
		configFileReader.MustHaveString("login.password"))

	log.Debug("STARTED")

	targets := configFileReader.MustHaveStringList("targets")
	targetCount := len(targets)

	if log.IsLevelEnabled(log.DebugLevel) {

		log.Debug("Count of targets: ", targetCount)

		for i := 0; i < targetCount; i++ {
			log.Debugln("Target: ", targets[i])
		}
	}

	for i := 0; i < targetCount; i++ {
		register(targets[i], login)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*port, nil)

	log.Debug("FINISHED")
}
