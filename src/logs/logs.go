package main

import (
	"github.com/sirupsen/logrus"
	"time"
	"os"
)

var log = logrus.New()

func main() {
	date := time.Now().Format("2006-01-02")

	file, err := os.OpenFile("src/logs/"+date+".log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	log.WithFields(logrus.Fields{
		"Log": "Log massages.",
	}).Info("Log massages.")
}
