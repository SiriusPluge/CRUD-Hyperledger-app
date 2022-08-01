package main

import (
	"app-hyperledger/internal/handlers"
	"app-hyperledger/internal/server"
	"app-hyperledger/internal/service"

	loggo "github.com/bukerdevid/log-go-log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loggo.InitCastomLogger(&logrus.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logrus.TraceLevel, false, true)

	if errConf := initConfig(); errConf != nil {
		logrus.Fatalf("error initializating configs: %s", errConf.Error())
	}

	service := service.NewService()
	handlers := handlers.NewHandler(service)

	srv := new(server.Server)
	if err := srv.RunServer(handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("pkg/configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
