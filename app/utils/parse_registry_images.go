package utils

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ParseRegistryImages(images string) []string {
	var imagesList []string
	if images == "" {
		if _, exists := os.LookupEnv("CI_REGISTRY_IMAGE"); !exists {
			log.Fatal("CI_REGISTRY_IMAGE is not defined.")
		} else {
			log.Warn("No Registry images defined. Will proceed with ", os.Getenv("CI_REGISTRY_IMAGE"))
			imagesList = append(imagesList, os.Getenv("CI_REGISTRY_IMAGE"))
			log.Infof("Retagging one image: %s", os.Getenv("CI_REGISTRY_IMAGE"))
			return imagesList
		}
	}

	log.Info("Processing images to retag.")
	imagesList = append(imagesList, strings.Fields(images)...)
	log.Infof("Images to retag: %v", imagesList)
	return imagesList
}
