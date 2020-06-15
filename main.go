package main

import (
	"docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"os"
	"time"
	log "github.com/sirupsen/logrus"

)
import "context"

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	client, err := docker.NewEnvClient()
	if err != nil {
		panic(err)
	}

	msgsCh, errCh := client.Events(context.Background(), types.EventsOptions{
		Since: time.Now().Add(-time.Hour).Format("2006-01-02T15:04:05"),
	})  // duplicate can easily be filtered outside of this api

	for {
		select {
		case msg := <- msgsCh:
			log.WithFields(log.Fields{
				"event": msg,
			}).Info("event")
		case err := <- errCh:
			log.WithFields(log.Fields{
				"error": err,
			}).Error("error")
			panic(err)
		}
	}
}

