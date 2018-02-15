//
// Copyright (c) 2017
// Cavium
// Mainflux
//
// SPDX-License-Identifier: Apache-2.0
//

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/edgexfoundry/edgex-go/support/logging"
)

const ()

// Set from the makefile
var version string = "undefined"

//var logger *zap.Logger

func main() {
	// logger, _ = zap.NewProduction()
	// defer logger.Sync()

	// distro.InitLogger(logger)

	// logger.Info("Starting export-distro", zap.String("version", version))
	cfg := loadConfig()

	errs := make(chan error, 2)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logging.StartHTTPServer(cfg, errs)

	c := <-errs
	fmt.Println("terminated: ", c)
	//logger.Info("terminated", zap.String("error", c.Error()))
}

func loadConfig() logging.Config {
	cfg := logging.GetDefaultConfig()
	return cfg
}

func env(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
