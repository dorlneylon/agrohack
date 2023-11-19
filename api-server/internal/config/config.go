package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string      `yaml:"env" env-default:"LOCAL"`
	HTTPServer  HTTPServer  `yaml:"http_server"`
	JsonStorage JsonStorage `yaml:"json_storage"`
	Service     Service     `yaml:"service"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"5s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

type JsonStorage struct {
	Path   string `yaml:"path"`
	Prefix string `yaml:"prefix"`
}

type Service struct {
	Command string `yaml:"command"`
	File    string `yaml:"file"`
}

type Flags struct {
	configPath string
}

func MustLoad() *Config {
	flags := parseFlags()
	if flags.configPath == "" {
		log.Fatal("no config path")
	}

	if _, err := os.Stat(flags.configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", flags.configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(flags.configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func parseFlags() *Flags {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config path")
	flag.Parse()
	return &Flags{
		configPath: configPath,
	}
}
