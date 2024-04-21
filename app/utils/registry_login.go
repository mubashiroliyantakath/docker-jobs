package utils

import (
	"os/exec"
	"strings"

	"github.com/mubashiroliyantakath/docker-jobs/app/models"
	log "github.com/sirupsen/logrus"
)

// Log into the Gitlab Container Registry for the project
func CIRegistryLogin() {

}

func LoginToRegistry(registry models.Registry) {
	// okBody, err := client.DockerClient.RegistryLogin(context.Background(), registry.AuthConfig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// registry.AuthConfig.IdentityToken = okBody.IdentityToken
	// log.Infof("Login successful: %v", okBody.Status)

	LogoutOfRegistry(registry)

	log.Infof("Logging into registry: %v", registry)
	loginCommand := []string{"docker", "login", "-u", registry.AuthConfig.Username, "-p", registry.AuthConfig.Password, registry.AuthConfig.ServerAddress}
	log.Debugf("Running command: %v", loginCommand)
	cmd := exec.Command(loginCommand[0], loginCommand[1:]...)
	li, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return
	}
	if strings.Contains(string(li), "Login Succeeded") {
		log.Infof("Login successful: %v", registry.AuthConfig.ServerAddress)
		return
	} else {
		log.Fatal("Login failed: ", string(li))
	}
}

func LogoutOfRegistry(registry models.Registry) {
	log.Infof("Cleaning up previous registry credentials: %v", registry.AuthConfig.ServerAddress)
	logoutCommand := []string{"docker", "logout", registry.AuthConfig.ServerAddress}
	cmd := exec.Command(logoutCommand[0], logoutCommand[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Info(string(output))
}
