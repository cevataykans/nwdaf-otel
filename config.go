package main

import (
	"nwdaf-otel/nrf"
	"nwdaf-otel/server"
)

var Config struct {
	A nwdaf.Config `yaml:"nwdaf"`
	B nrf.Config   `yaml:"nrf"`
}
