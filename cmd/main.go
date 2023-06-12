package main

import (
	"github.com/Phaseant/DreamAnalyzer/internal/handler"
	"github.com/Phaseant/DreamAnalyzer/internal/repository"
	"github.com/Phaseant/DreamAnalyzer/internal/server"
	"github.com/Phaseant/DreamAnalyzer/internal/service"
	"github.com/xlab/closer"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	setLogger()
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	gptClient := repository.NewGptClient(viper.GetString("TOKEN"))
	repo := repository.NewRepository(gptClient)
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	go func() {
		srv := new(server.Server)
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("Error while running server: %v", err)
		}
		closer.Close()
	}()

	closer.Hold()

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func setLogger() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03",
	})
}
