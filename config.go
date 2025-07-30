package main

import (
	"nwdaf-otel/clients/nrf"
)

var Config struct {
	A nwdaf.Config `yaml:"nwdaf"`
	B nrf.Config   `yaml:"nrf"`
}
