package pkg

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"strings"
)

type Config struct {
	Redis    Redis    `yaml:"redis"`
	Postgres Postgres `yaml:"postgres_master"`
	HttpServer
}

type HttpServer struct {
	Address string `yaml:"address"`
}

type Postgres struct {
	Host             string `yaml:"host"`
	Port             int    `yaml:"port"`
	DB               string `yaml:"db"`
	User             string `yaml:"user"`
	Pass             string `yaml:"pass"`
	BatchCount       int    `yaml:"batch_count"`
	ConnectionsCount int    `yaml:"connections_count"`
}

type Redis struct {
	Enable     bool   `yaml:"enable"`
	Address    string `yaml:"address"`
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	DB         int    `yaml:"db"`
	Pass       string `yaml:"pass"`
	PoolSize   int    `yaml:"pool_size"`
	MaxRetries int    `yaml:"max_retries"`
}

func NewConfig(serviceName string, path string) *Config {
	conf := &Config{}
	err := conf.loadConf(serviceName, path)
	if err != nil {
		log.Fatalf("Load yaml config file error: '%v'", err)
		return nil
	}
	return conf
}

func (conf *Config) configureViper(serviceName string) {
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix(serviceName)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/." + serviceName)
	viper.AddConfigPath("/etc/" + serviceName + "/")
	viper.SetConfigName("config")
}

func (conf *Config) loadConfFromFile(path string) error {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("File does not exist : ", path)
		return err
	}

	if err := viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
		return err
	}

	return nil
}

func (conf *Config) loadConf(serviceName string, path string) error {
	conf.configureViper(serviceName)
	err := conf.loadConfFromFile(path)
	if err != nil {
		return err
	}
	conf.Reload()
	return nil
}

func (conf *Config) Reload() {
	// Redis
	conf.Redis.Address = viper.GetString("redis.address")
	conf.Redis.Host = viper.GetString("redis.host")
	conf.Redis.Port = viper.GetString("redis.port")
	conf.Redis.DB = viper.GetInt("redis.db")
	conf.Redis.Pass = viper.GetString("redis.pass")
	conf.Redis.PoolSize = viper.GetInt("redis.pool_size")
	conf.Redis.MaxRetries = viper.GetInt("redis.max_retries")
	conf.Redis.Enable = viper.GetBool("redis.enable")

	// Postgres
	conf.Postgres.Host = viper.GetString("postgres.host")
	conf.Postgres.Port = viper.GetInt("postgres.port")
	conf.Postgres.DB = viper.GetString("postgres.db")
	conf.Postgres.User = viper.GetString("postgres.user")
	conf.Postgres.Pass = viper.GetString("postgres.pass")
	conf.Postgres.ConnectionsCount = viper.GetInt("postgres.connections_count")

}
