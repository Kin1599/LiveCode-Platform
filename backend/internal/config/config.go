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

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}

func ConStringFromCfg(storageCfg StorageConfig) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		storageCfg.User,
		storageCfg.Pass,
		storageCfg.Host,
		storageCfg.Port,
		storageCfg.Name,
	)
}
