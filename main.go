package main

import (
	"flag"
	"github.com/spf13/viper"
	"go-framework/http"
	"go-framework/mysql"
	"log"
)

var (
	configPath = flag.String("config_path", "", "config path")
)

func configInit() {
	if *configPath == "" {
		log.Fatal("config path is nil")
	}
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(*configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("read config failed: %v", err)
	}
}

func main() {
	flag.Parse()
	configInit()

	mysql.Init()

	err := http.Init()
	if err != nil {
		log.Fatal(err)
	}
}
