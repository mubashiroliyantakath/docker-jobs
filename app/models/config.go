package models

import (
	"fmt"

	"github.com/docker/docker/api/types/registry"
)

type Registry struct {
	AuthConfig registry.AuthConfig
	Enabled    bool
}

type Metadata struct {
	Builder    string   `mapstructure:"builder"`
	Version    string   `mapstructure:"version"`
	Tags       []string `mapstructure:"tags"`
	Images     string   `mapstructure:"images"`
	Registries string   `mapstructure:"registries"`
}

func (r Registry) String() string {
	return fmt.Sprintf(r.AuthConfig.ServerAddress)
}
