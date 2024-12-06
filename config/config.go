package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type NoConfig struct {
}

func (NoConfig) Error() string {
	return "no config found"
}

type Config struct {
	SchemaSQL  string `yaml:"schemasql"`
	Transforms string `yaml:"transforms"`
}

func LoadConfig() (*Config, error) {
	var config Config
	data, err := readConfig()
	if err != nil {
		return &config, err
	}
	yaml.Unmarshal(data, &config)
}

func readConfig() ([]byte, error) {
	cFile := findConfig()
	var bytes []byte
	if cFile == "" {
		return bytes, &NoConfig{}
	}
	return os.ReadFile(cFile)
}

func findConfig() string {
	homeDir, _ := os.UserHomeDir()

	paths := []string{"config.yaml", "../config.yaml", homeDir + "/.dough/config.yaml"}
	cnfPath := ""
	for _, p := range paths {

		_, err := os.Stat(p)
		if err == nil {
			cnfPath = p
			break
		}
	}

	return cnfPath

}
