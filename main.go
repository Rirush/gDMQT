package main

import (
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("TOML")
	if runtime.GOOS == "linux" {
		dir, err := filepath.Abs(filepath.Dir(os.Args[0])) // XXX: Unreliable
		if err != nil {                                    // XXX: unexpected
			panic(err)
		}
		if strings.Contains(dir, "bin") {
			viper.SetDefault("Log.file", "/var/log/mqtt/broker.log")
			viper.AddConfigPath("/etc/gomqtt/")
		} else {
			viper.AddConfigPath(".")
		}
		// TODO: viper.SetDefault("Log.system", false)
	} else {
		viper.AddConfigPath(".")
		viper.SetDefault("Log.file", "./broker.log")
	}
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.New("failed to read config file: " + err.Error()))
	}

	logFilePath := viper.GetString("Log.file")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
	if !(logFilePath == "" || !strings.ContainsAny(logFilePath, "/\\")) {
		logFile, err := os.OpenFile(logFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0664)
		if err != nil {
			panic(errors.New("failed to open log file: " + err.Error()))
		}
		multiWriter := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(multiWriter)
	}

	log.Println("GoMQTT Broker")
	log.Println("Copyright © 2019 Vladyslav Yamkovyi (Hexawolf)")
}
