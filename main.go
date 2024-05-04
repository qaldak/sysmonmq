package main

import (
	"flag"
	"log"

	//"os"
	//"os/signal"
	//"syscall"

	"github.com/joho/godotenv"

	collector "github.com/qaldak/SysMonMQ/systeminfo"
	logger "github.com/qaldak/SysMonMQ/logging"
	// mqtt "github.com/qaldak/SysMonMQ/mqtt"
)

func main() {
	// TODO: logger
	logger.Info("Foo")
	
	si, err := collector.GetSystemInfo()
	if err != nil {
		logger.Error("Error while get systeminfo: ", err)
	}
	logger.Debug(si)

	/* JSON structure
	JSON_structure='{"CPU01":'$CPU01',"CPU05":'$CPU05',"CPU15":'$CPU15',"CPU_temp":'$CPU_temp',"RAM_total":'$RAM_total',"RAM_free":'$RAM_free',"RAM_avlbl":'$RAM_avlbl',"RAM_used":'$RAM_used',"Disk_total":'$Disk_total',"Disk_free":'$Disk_free',"Disk_used":'$Disk_used',"Sys_Uptime":'$Sys_uptime',"LastLogin_date":"'$LastLogin_date'","LastLogin_user":"'$LastLogin_user'","LastLogin_from":"'$LastLogin_from'"}'
	*/ 


	// TODO: generate json

	// TODO: mqtt message to broker

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

	// TODO: mqtt.InitMQTT()
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
