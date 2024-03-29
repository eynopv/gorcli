package internal

import (
	"encoding/json"
	"fmt"

	"github.com/eynopv/gorcli/internal/utils"
)

type Config struct {
	ShowHeaders bool              `json:"showHeaders"`
	Headers     map[string]string `json:"headers"`
	Variables   map[string]string `json:"variables"`
}

func LoadConfig() (*Config, error) {
	var config Config

	configPath := "./gorcli.config.json"

	if content := utils.LoadFile(configPath); content != nil {
		err := json.Unmarshal(*content, &config)
		if err != nil {
			fmt.Println("Failed to load config file:", err)
			return nil, err
		}

		// TODO: headers could use variables so maybe params should be parsed right before request
		for key, value := range config.Headers {
			param := Param{Name: key, Value: value}
			config.Headers[key] = param.ParseValue()
		}

		return &config, nil
	}

	config = Config{
		Headers:   map[string]string{},
		Variables: map[string]string{},
	}

	return &config, nil
}
