package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var ProjectRootPath = ConfigsDirPath() + "/../"

type Env int

const (
	Development Env = iota
	Production
)

var CurrentEnv Env = Development

type (
	Config struct {
		Grpc struct {
			Host string `koanf:"host"`
			Port int    `koanf:"port"`
		} `koanf:"grpc"`
	}
)

func ConfigsDirPath() string {
	_, f, _, ok := runtime.Caller(0)
	if !ok {
		panic("Error in generating env dir")
	}

	return filepath.Dir(f)
}

func Read() *Config {
	var fileName string

	// Load KAVKA ENV
	env := strings.ToLower(os.Getenv("KAVKA_ENV"))

	if len(strings.TrimSpace(env)) == 0 || env == "development" {
		CurrentEnv = Development
		fileName = "config.development.yml"
	} else if env == "production" {
		CurrentEnv = Production
		fileName = "config.production.yml"
	} else {
		log.Fatalln(errors.New("Invalid env value set for variable KAVKA_ENV: " + env))
	}

	// Load YAML configs
	k := koanf.New(ConfigsDirPath())
	if err := k.Load(file.Provider(fmt.Sprintf("%s/%s", ConfigsDirPath(), fileName)), yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	config := &Config{}
	if err := k.Unmarshal("", config); err != nil {
		log.Fatalf("error unmarshaling config: %v", err)
	}

	return config
}
