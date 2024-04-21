package main

import (
	"os"

	"github.com/mubashiroliyantakath/docker-jobs/app/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetOutput(os.Stdout)
	logLevel, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)
	log.Debug("Log level set to: ", logLevel)

	err = utils.NewAppConfig()
	if err != nil {
		log.Fatal("Failed to load config: ", err)
	}
}

func main() {
	// This is a comment

	images := utils.ParseRegistryImages(utils.AppConfig.Images)
	registries := utils.ParseRegistries(utils.AppConfig.Registries)

	if len(registries) == 0 {
		log.Fatal("No new registries defined.")
	}

	for _, registry := range registries {
		utils.LoginToRegistry(registry)
	}

	for _, image := range images {
		_ = utils.RetagImage(image, utils.AppConfig.Tags, registries)
		// tag image to the ones in the returned list

	}

	defer func() {
		for _, registry := range registries {
			utils.LogoutOfRegistry(registry)
		}
	}()

}
