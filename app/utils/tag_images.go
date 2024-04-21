package utils

import (
	"os"
	"strings"

	"github.com/mubashiroliyantakath/docker-jobs/app/models"
	log "github.com/sirupsen/logrus"
)

func RetagImage(image string, tags []string, registries []models.Registry) []string {
	var imageList []string
	for _, registry := range registries {
		newImage := strings.ReplaceAll(image, os.Getenv("CI_REGISTRY"), registry.AuthConfig.ServerAddress)
		for _, tag := range tags {
			imageList = append(imageList, newImage+":"+tag)
		}
	}
	log.Info("Retag image list for ", image, " => ", imageList)
	return imageList
}
