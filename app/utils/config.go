package utils

import (
	"strings"

	"github.com/mubashiroliyantakath/docker-jobs/app/models"
	"github.com/spf13/viper"
)

var (
	AppConfig *models.Metadata
)

func NewAppConfig() error {

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigType("yaml")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	v.SetConfigFile("docker-build-config.yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(&AppConfig); err != nil {
		return err
	}

	return nil
}
