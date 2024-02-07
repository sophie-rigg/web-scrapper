package main

import (
	"flag"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sophie-rigg/web-scrapper/logging"
	"github.com/sophie-rigg/web-scrapper/processor/node"
	"github.com/sophie-rigg/web-scrapper/reader/http"
	"github.com/sophie-rigg/web-scrapper/worker"
)

var (
	logLevel logging.LogLevel
	domain   string
)

func init() {
	flag.Var(&logLevel, "level", "log level")
	flag.StringVar(&domain, "domain", "https://www.monzo.com", "domain to scrape")
}

func main() {
	flag.Parse()

	logrus.Info("Starting web-scrapper")

	reader := http.New()

	domain = strings.TrimSuffix(domain, "/")

	domainProcessor := node.New(domain)

	worker.New(domainProcessor, reader, domain).Do()
}
