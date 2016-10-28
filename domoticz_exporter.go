// Copyright (C) 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/log"

	"github.com/nlamirault/domoticz_exporter/domoticz"
	exporter_version "github.com/nlamirault/domoticz_exporter/version"
)

const (
	banner = "domoticz_exporter - %s\n"

	namespace = "domoticz"
)

var (
	debug         bool
	version       bool
	listenAddress string
	metricsPath   string
	domoticzPath  string

	last = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "last"),
		"Last Push from Domoticz.",
		nil, nil,
	)
)

// Exporter collects Domoticz stats from the given server and exports them using
// the prometheus metrics package.
type Exporter struct {
}

// NewExporter returns an initialized Exporter.
func NewExporter() (*Exporter, error) {
	log.Infoln("Setup Domoticz exporter")
	return &Exporter{}, nil
}

// Describe describes all the metrics ever exported by the Domoticz exporter.
// It implements prometheus.Collector.
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- last
}

// Collect the stats from channel and delivers them as Prometheus metrics.
// It implements prometheus.Collector.
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	log.Infof("Domoticz exporter starting")

	ch <- prometheus.MustNewConstMetric(
		last, prometheus.GaugeValue, float64(time.Now().UnixNano()/1e3),
	)

	log.Infof("Domoticz exporter finished")
}

func (e *Exporter) DomoticzHandler(w http.ResponseWriter, r *http.Request) {
	log.Debugf("Receive Domoticz post")
	var metric domoticz.Metric
	err := json.NewDecoder(r.Body).Decode(&metric)
	if err != nil {
		log.Errorf("Error Decoding Domoticz request: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Debugf("Domoticz metric received: %s", metric)
}

func init() {
	// parse flags
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.StringVar(&listenAddress, "web.listen-address", ":9112", "Address to listen on for web interface and telemetry.")
	flag.StringVar(&metricsPath, "web.telemetry-path", "/metrics", "Path under which to expose metrics.")
	flag.StringVar(&domoticzPath, "domoticz.path", "/domoticz", "Path to accept POST requests from domoticz.")

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, fmt.Sprintf(banner, exporter_version.Version))
		flag.PrintDefaults()
	}

	flag.Parse()

	if version {
		fmt.Printf("%s", exporter_version.Version)
		os.Exit(0)
	}
}

func main() {
	exporter, err := NewExporter()
	if err != nil {
		log.Errorf("Can't create exporter : %s", err)
		os.Exit(1)
	}
	log.Infoln("Register exporter")
	prometheus.MustRegister(exporter)

	http.Handle(metricsPath, prometheus.Handler())
	http.HandleFunc(domoticzPath, exporter.DomoticzHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
             <head><title>Domoticz Exporter</title></head>
             <body>
             <h1>Domoticz Exporter</h1>
             <p><a href='` + metricsPath + `'>Metrics</a></p>
             </body>
             </html>`))
	})

	log.Infoln("Listening on", listenAddress)
	log.Fatal(http.ListenAndServe(listenAddress, nil))
}
