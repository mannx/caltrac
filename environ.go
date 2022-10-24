package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

//
//	This file contains code for managing static paths that might change with environment variables
//	and other data to prevent hard coding in multiple places

type EnvironmentDefinition struct {
	DataPath string `envconfig:"CALTRAC_DATA_PATH"`

	// UserID and GroupID are used to set the file permissions for all exported files
	UserID  int `envconfig:"PUID"` // userid the container should be running under
	GroupID int `envconfig:"PGID"` // groupid "								"
}

var Environment = EnvironmentDefinition{}

func (e *EnvironmentDefinition) Init() {
	e.Default()

	err := envconfig.Process("", e)
	if err != nil {
		log.Error().Err(err).Msg("Unable to parse local environment")
		return
	}
}

func (e *EnvironmentDefinition) Default() {
	e.DataPath = "/data"
}
