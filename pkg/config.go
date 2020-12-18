package pkg

import (
	"bytes"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type Config struct {
	Redis         Redis         `yaml:"redis"`
	Postgres      Postgres      `yaml:"postgres_master"`
	HttpServer    HttpServer    `yaml:"http_server"`
	PromotionCode PromotionCode `yaml:"promotion_code"`
	EndPoints     EndPoints     `yaml:"end_points"`
	Kafka         Kafka         `yaml:"kafka"`
}

type Kafka struct {
	Enable           bool   `yaml:"enable"`
	BootstrapServers string `yaml:"bootstrap_servers"`
	GroupID          string `yaml:"group_id"`
	AutoOffsetReset  string `yaml:"auto_offset_reset"`
	Topic            string `yaml:"topic"`
}

type EndPoints struct {
	Wallet string `yaml:"wallet"`
}

type PromotionCode struct {
	CodeValue string    `yaml:"code_value"`
	LockKey   string    `yaml:"lock_key"`
	StartTime time.Time `yaml:"start_time"`
	EndTime   time.Time `yaml:"end_time"`
	MaxCodes  int32     `yaml:"max_codes"`
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

	// HttpServer
	conf.HttpServer.Address = viper.GetString("http_server.address")

	// PromotionCode
	conf.PromotionCode.CodeValue = viper.GetString("promotion_code.code_value")
	conf.PromotionCode.LockKey = viper.GetString("promotion_code.lock_key")
	conf.PromotionCode.MaxCodes = viper.GetInt32("promotion_code.max_codes")
	conf.PromotionCode.StartTime = viper.GetTime("promotion_code.start_time")
	conf.PromotionCode.EndTime = viper.GetTime("promotion_code.end_time")

	// EndPoints
	conf.EndPoints.Wallet = viper.GetString("end_points.wallet")

	// Kafka
	conf.Kafka.BootstrapServers = viper.GetString("kafka.bootstrap_servers")
	conf.Kafka.Topic = viper.GetString("kafka.topic")
	conf.Kafka.Enable = viper.GetBool("kafka.enable")
	conf.Kafka.GroupID = viper.GetString("kafka.group_id")
	conf.Kafka.AutoOffsetReset = viper.GetString("kafka.auto_offset_reset")
}
