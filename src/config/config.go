package config

import (
	"errors"
	"flag"
	"os"

	"github.com/paingha/pex/src/controllers"
	"github.com/paingha/pex/src/lib"
	"github.com/paingha/pex/src/models"
	"gopkg.in/yaml.v2"
)

var (
	// ErrIsADir is returned when the provided file path is a directory not a file.
	ErrIsADir = errors.New("path is a directory not a file")
)

type SystemConfig struct {
	Environment string `yaml:"env"`
	Port        string `yaml:"port"`
}

func NewSystemConfig() *SystemConfig {
	return &SystemConfig{}
}

func NewENV(db *models.FibonacciDataStore, logger *lib.Log) *controllers.Env {
	return &controllers.Env{
		Db:     db,
		Logger: logger,
	}
}

func InitSystemConfig(path string) (*SystemConfig, error) {
	if path == "" {
		configFilePath, err := parseFlags()
		if err != nil {
			return nil, err
		}
		path = configFilePath
	}
	systemConfig := NewSystemConfig()
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	yamlDecoder := yaml.NewDecoder(file)
	if err := yamlDecoder.Decode(&systemConfig); err != nil {
		return nil, err
	}
	return systemConfig, nil
}

// parseFlags will create and parse the CLI flags
func parseFlags() (string, error) {
	configPath := ""
	flag.StringVar(&configPath, "config", "./config.yaml", "configuration file path")
	flag.Parse()

	if err := validateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}

// validateConfigPath validates the configpath.
func validateConfigPath(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return err
	}
	if info.IsDir() {
		return ErrIsADir
	}
	return nil
}
