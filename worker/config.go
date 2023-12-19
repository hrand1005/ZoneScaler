package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	CoordinatorHost string `json:"coordinator_host"`
	CoordinatorPort int    `json:"coordinator_port"`
}

// LoadConfig parses worker configuration data from a JSON file.
func LoadConfig(filename string) (Config, error) {
	var config Config

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, fmt.Errorf("failed to read config file: %v", err)
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	return config, nil
}
