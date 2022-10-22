package main

import (
	"context"
	"encoding/json"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/buley/n3ut/interfaces"
	"github.com/buley/n3ut/interfaces/server"
)

const (
	DefaultConfigPath = "_config/sample.json"
)

func loadConfig(path string) (*server.Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var conf server.Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, err
	}
	if err := conf.Validate(); err != nil {
		return nil, err
	}

	return &conf, nil
}

func main() {
	var confPath = flag.String("conf", DefaultConfigPath, "path to your config.json")
	flag.Parse()

	conf, err := loadConfig(*confPath)
	if err != nil {
		panic(err)
	}

	srv := interfaces.NewServer(conf)

	done := make(chan bool, 1)
	go func() {
		sigterm := make(chan os.Signal, 1)
		signal.Notify(sigterm, syscall.SIGTERM)
		<-sigterm

		if err := srv.Shutdown(context.Background()); err != nil {
			srv.Logger().Fatal(err)
		}
		close(done)
	}()
	if err := srv.Start(); err != nil {
		srv.Logger().Info(err)
	}
	<-done
}
