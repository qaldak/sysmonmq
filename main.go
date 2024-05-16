package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/joho/godotenv"

	logger "github.com/qaldak/sysmonmq/internal/logging"
	"github.com/qaldak/sysmonmq/internal/utils"
	"github.com/qaldak/sysmonmq/mqtt"
	"github.com/qaldak/sysmonmq/slack"
	collector "github.com/qaldak/sysmonmq/systeminfo"
)

func main() {
	// TODO: logger
	logger.Info("Foo")

	// collect system information
	data, err := collector.GetSystemInfo()
	if err != nil {
		logger.Error("Error while get systeminfo: ", err)
		
	}
	logger.Info("Collected system infos: ", data)

	// generate json
	json, err := utils.GenerateJson(data)
	if err != nil {
		logger.Error("Failed to generate JSON")
	}
	logger.Info("JSON: ", json)

	// pubslish mqtt message to broker
	err = mqtt.PublishMessage(json)
	if err != nil {
		logger.Error("Failed to publish MQTT message")
		slack.PostSlackMsg(fmt.Sprint(err))
	}

	log.Print("Foo")
}

func init() {
	// input args declaration
	var debug bool

	// flags declaration usig flag package
	flag.CommandLine.BoolVar(&debug, "debug", false, "--debug: set loglevel to 'debug'")
	flag.Parse() // after declaring flags, we need to call it

	if isFlagPassed("debug") {
		debug = true
	}

	// initialize logger
	logger.InitLogger(debug)
	logger.Info("Yeah, logger is up an running", "Foo")

	// load env settings
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error("Error initializing environment variables from '.env': ", err)
	}
}

func isFlagPassed(name string) bool {
	flagFound := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			flagFound = true
		}
	})
	return flagFound
}
