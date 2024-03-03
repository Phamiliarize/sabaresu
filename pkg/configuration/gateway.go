package configuration

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Route struct {
	Method    string
	Path      string
	Functions []string
}

type GatewayConfiguration []Route

func LoadGatewayConfiguration() (*GatewayConfiguration, error) {
	var cfg map[string]GatewayConfiguration
	var routes GatewayConfiguration

	tomlCfg, err := os.ReadFile("./gateway.toml")
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal([]byte(tomlCfg), &cfg)
	if err != nil {
		return nil, err
	}

	routes = cfg["routes"]

	return &routes, nil
}
