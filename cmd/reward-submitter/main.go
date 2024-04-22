package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"

	"github.com/metisprotocol/metis-reward-submitter/internal/config"
	"github.com/metisprotocol/metis-reward-submitter/internal/submitter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	var (
		configPath string
		debug      bool
	)
	flag.StringVar(&configPath, "conf", "config.yaml", "the config file path")
	flag.BoolVar(&debug, "debug", false, "debug mode")
	flag.Parse()

	basectx, cancel := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if debug {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	slog.Info("parsing config file", "path", configPath)
	conf, err := config.Parse(configPath)
	if err != nil {
		slog.Error("failed to parse the config", "err", err)
		return
	}

	reg := prometheus.NewRegistry()
	metric := submitter.NewMetrics(reg)

	go func() {
		server := &http.Server{Addr: ":9090"}
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
		_ = server.ListenAndServe()
	}()

	srv, err := submitter.New(basectx, metric, conf)
	if err != nil {
		slog.Error("failed to init submitter", "err", err)
	}
	defer srv.Stop()

	srv.Start(basectx)
}
