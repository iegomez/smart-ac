package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/BurntSushi/toml"
	"github.com/iegomez/smart-ac/internal/api"
	"github.com/iegomez/smart-ac/internal/config"
	"github.com/iegomez/smart-ac/internal/storage"
)

func main() {

	confFile := flag.String("conf", "smart-ac.toml", "path to toml configuration file")
	flag.Parse()

	err := importConf(*confFile)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	tasks := []func() error{
		setLogLevel,
		setupStorage,
		setupAPI,
	}

	for _, t := range tasks {
		if err := t(); err != nil {
			log.Fatal(err)
		}
	}

	sigChan := make(chan os.Signal)
	exitChan := make(chan struct{})
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	log.WithField("signal", <-sigChan).Info("signal received")
	go func() {
		log.Warning("stopping smart-ac")
		exitChan <- struct{}{}
	}()
	select {
	case <-exitChan:
	case s := <-sigChan:
		log.WithField("signal", s).Info("signal received, stopping immediately")
	}

}

func importConf(filename string) error {
	log.Infoln("parsing ", filename)
	if _, err := toml.DecodeFile(filename, &config.C); err != nil {
		return err
	}
	log.Infof("config: %#v\n", config.C)
	return nil
}

func setLogLevel() error {
	log.Infoln("setting log level to ", config.C.General.LogLevel)
	log.SetLevel(log.Level(uint8(config.C.General.LogLevel)))
	return nil
}

func setupStorage() error {
	log.Infoln("setting up storage")
	if err := storage.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup storage error")
	}

	return nil
}

func setupAPI() error {
	log.Infoln("setting up API")
	if err := api.Setup(config.C); err != nil {
		return errors.Wrap(err, "setup api error")
	}
	return nil
}
