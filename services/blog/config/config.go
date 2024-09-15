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
	Test
)

var CurrentEnv Env = Development

func (c Env) String() string {
	if c == Development {
		return "development"
	} else if c == Production {
		return "production"
	} else if c == Test {
		return "test"
	} else {
		return "UNKNOWN"
	}
}

type (
	Config struct {
		Grpc struct {
			Host           string `koanf:"host"`
			Port           int    `koanf:"port"`
			AuthServiceUrl string `koanf:"auth_service_url"`
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

	// Load GO_ENV
	env := strings.ToLower(os.Getenv("GO_ENV"))

	if len(strings.TrimSpace(env)) == 0 || env == "development" {
		CurrentEnv = Development
		fileName = "config.development.yml"
	} else if env == "production" {
		CurrentEnv = Production
		fileName = "config.production.yml"
	} else if env == "test" {
		CurrentEnv = Test
		fileName = "config.test.yml"
	} else {
		log.Fatalln(errors.New("Invalid env value set for variable GO_ENV: " + env))
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
