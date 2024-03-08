package main

import (
	"context"
	"flag"
	"log/slog"
	"os/signal"
	"syscall"

	"github.com/metisprotocol/metis-seq-reward-submitter/internal/config"
	"github.com/metisprotocol/metis-seq-reward-submitter/internal/submitter"
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

	srv, err := submitter.New(basectx, conf)
	if err != nil {
		slog.Error("failed to init submitter", "err", err)
	}
	defer srv.Stop()

	srv.Start(basectx)
}
