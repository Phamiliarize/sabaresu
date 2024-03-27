package gateway

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Route struct {
	Method      string
	Path        string
	InputSchema string
	FuncDir     string
	Functions   []string
}

type Gateway struct {
	Routes []Route
	Schema map[string]interface{}
}

func NewGatewayFromCfg(cfgPath *string, schemaPath *string) (*Gateway, error) {
	var cfg map[string][]Route

	// Load Gateway Configuration
	tomlCfg, err := os.ReadFile(*cfgPath)
	if err != nil {
		return nil, err
	}

	err = toml.Unmarshal([]byte(tomlCfg), &cfg)
	if err != nil {
		return nil, err
	}

	gateway := &Gateway{
		Routes: cfg["routes"],
		Schema: map[string]interface{}{},
	}

	// Load all Schema
	dir, err := os.Open(*schemaPath)
	if err != nil {
		fmt.Println("Error opening schema directory:", err)
		return nil, err
	}
	defer dir.Close()

	fileNames, err := dir.Readdirnames(0)
	if err != nil {
		fmt.Println("Error reading files from directory:", err)
		return nil, err
	}

	for _, fileName := range fileNames {
		if filepath.Ext(fileName) == ".toml" {
			filePath := filepath.Join(*schemaPath, fileName)
			content, err := os.ReadFile(filePath)
			if err != nil {
				return nil, err
			}

			var raw map[string]interface{}
			_, err = toml.Decode(string(content), &raw)
			if err != nil {
				return nil, err
			}

			gateway.Schema[strings.TrimSuffix(fileName, filepath.Ext(fileName))] = raw

		}
	}

	return gateway, nil
}
