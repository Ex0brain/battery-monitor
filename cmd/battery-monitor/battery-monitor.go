package main

import (
	"github.com/distatus/battery"
	"github.com/maateen/battery-monitor/internal/monitor"
	"github.com/maateen/battery-monitor/internal/notification"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	batteries, err := battery.GetAll()

	if err != nil {
		log.Fatal(err)
	} else if len(batteries) == 0 {
		log.Fatal("No battery found.")
	} else if len(batteries) == 1 {
		log.Info("Battery found.")
	} else {
		log.Warn("Multiple batteries found but only one supported.")
	}

	messageChannel := make(chan monitor.Message)
	go monitor.Monitor(messageChannel)

	for {
		select {
		case message := <-messageChannel:
			notification.Notify(message)
		}
	}
}