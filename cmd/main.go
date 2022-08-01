package main

import (
	"app-hyperledger/internal/bcapi"
	"app-hyperledger/internal/handlers"
	"app-hyperledger/internal/server"
	"app-hyperledger/internal/service"
	"app-hyperledger/pkg/models"

	loggo "github.com/bukerdevid/log-go-log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	loggo.InitCastomLogger(&logrus.JSONFormatter{TimestampFormat: "15:04:05 02/01/2006"}, logrus.TraceLevel, false, true)

	if errConf := initConfig(); errConf != nil {
		logrus.Fatalf("error initializating configs: %s", errConf.Error())
	}

	cfg := getConfig()

	contract, err := bcapi.Init(cfg)
	if err != nil {
		logrus.Fatalf("error initializating connect to BC: %s", err.Error())
	}

	service := service.NewService(contract)
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

func getConfig() *models.ConfigBC {
	cfg := &models.ConfigBC{
		TlsCertPath:   viper.GetString("ledger.tlsCertPath"),
		PeerEndpoint:  viper.GetString("ledger.peerEndpoint"),
		GatewayPeer:   viper.GetString("ledger.gatewayPeer"),
		ChannelName:   viper.GetString("ledger.channelName"),
		ChaincodeName: viper.GetString("ledger.chaincodeName"),
		UserCert: models.UserCert{
			OrgID:   viper.GetString("ledger.mspID"),
			Cert:    viper.GetString("ledger.certPath"),
			PrivKey: viper.GetString("ledger.keyPath"),
		},
	}

	logrus.Info(cfg)

	return cfg
}
