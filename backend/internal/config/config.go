package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-default:"local"`
	StoragePath StorageConfig `yaml:"storage_path"`
}

type StorageConfig struct {
	Host       string `yaml:"db_host"`
	Name       string `yaml:"db_name"`
	User       string `yaml:"db_user"`
	Pass       string `yaml:"db_pass"`
	Port       int    `yaml:"db_port"`
	BucketName string `yaml:"bucket_name"`
}

func Load() (*Config, error) {
	configPath, err := fetchConfigPath()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch config path: %w", err)
	}

	if configPath == "" {
		return nil, fmt.Errorf("config path is empty")
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, fmt.Errorf("failed to read config from %s: %w", configPath, err)
	}

	return &cfg, nil
}

func MustLoad() *Config {
	cfg, err := Load()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}
	return cfg
}

func fetchConfigPath() (string, error) {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	if configPath == "" {
		configPath = os.Getenv("CONFIG_PATH")
	}

	return configPath, nil
}

func BuildDBConnectionString(cfg StorageConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Pass,
		cfg.Host,
		cfg.Port,
		cfg.Name,
	)
}
