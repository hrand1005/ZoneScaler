package coordinator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Host      string `json:"host"`
	Port      int    `json:"port"`
	StaticDir string `json:"static_dir"`
}

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
