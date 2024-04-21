package utils

import (
	"strings"

	"github.com/docker/docker/api/types/registry"
	log "github.com/sirupsen/logrus"
	"gitlab.com/devletix/devops/docker-jobs/app/models"
)

func ParseRegistries(repos string) []models.Registry {

	var registries []models.Registry
	uniqueServerAddresses := make(map[string]struct{})

	repoList := strings.Fields(repos)
	for counter, repo := range repoList {
		var repository string
		var username string
		var password string
		enabled := true

		inputMap := make(map[string]string)
		pairs := strings.Split(repo, ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 {
				inputMap[kv[0]] = kv[1]
			}
		}

		if e, ok := inputMap["enabled"]; ok {
			if e == "false" {
				log.Infof("Registry is disabled in input %d, skipping", counter+1)
				continue
			}
		}

		if r, ok := inputMap["registry"]; ok {
			repository = r
		} else {
			log.Warnf("Registry is not defined for input %d, skipping", counter+1)
			continue
		}

		if u, ok := inputMap["username"]; ok {
			username = u
		} else {
			log.Warnf("Username not defined for input %d, skipping", counter+1)
			continue
		}

		if p, ok := inputMap["password"]; ok {
			password = p
		} else {
			log.Warnf("Password not defined for input %d, skipping", counter+1)
			continue
		}
		reg := models.Registry{
			AuthConfig: registry.AuthConfig{
				Username:      username,
				Password:      password,
				ServerAddress: repository,
			},
			Enabled: enabled,
		}

		if _, exists := uniqueServerAddresses[reg.AuthConfig.ServerAddress]; exists {
			log.Warnf("Credentials already parsed for registry host '%s', skipping input %d", reg.AuthConfig.ServerAddress, counter+1)
			continue
		} else {
			uniqueServerAddresses[reg.AuthConfig.ServerAddress] = struct{}{}
		}
		registries = append(registries, reg)
		log.Debugf("Registry added from input %d: %v", counter+1, reg)

	}
	log.Info("Registries parsed: ", registries)
	return registries
}
