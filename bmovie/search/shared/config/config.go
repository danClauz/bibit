package config

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
)

var (
	envVar     = flag.String("env", "local", "Application environment")
	configJson = flag.String("config_filepath", "./config-local.json", "Configuration file to load")
)

type (
	server struct {
		Host         string        `json:"host"`
		Port         string        `json:"port"`
		ReadTimeout  time.Duration `json:"read_timeout"`
		WriteTimeout time.Duration `json:"write_timeout"`
	}
	mysql struct {
		Master *database `json:"master"`
		Slave  *database `json:"slave"`
	}
	database struct {
		DSN             string        `json:"dsn"`
		MaxIdleConn     int           `json:"max_idle_conn"`
		MaxOpenConn     int           `json:"max_open_conn"`
		ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	}
	logger struct {
		FilePath     string `json:"file_path"`
		FileName     string `json:"file_name"`
		Formatter    string `json:"formatter"`
		Stdout       bool   `json:"stdout"`
		ReportCaller bool   `json:"report_caller"`
	}
	datadog struct {
		Host      string `json:"host"`
		Port      string `json:"port"`
		Namespace string `json:"namespace"`
		IsEnabled bool   `json:"is_enabled"`
	}
	newRelic struct {
		AppName   string `json:"app_name"`
		AppKey    string `json:"app_key"`
		IsEnabled bool   `json:"is_enabled"`
	}

	Omdb struct {
		Host string `json:"host"`
		Key  string `json:"key"`
	}

	EnvConfig struct {
		env        string
		HttpServer *server   `json:"http_server"`
		GrpcServer *server   `json:"grpc_server"`
		Mysql      *mysql    `json:"mysql"`
		Logger     *logger   `json:"logger"`
		Datadog    *datadog  `json:"datadog"`
		NewRelic   *newRelic `json:"new_relic"`
		Omdb       *Omdb     `json:"omdb"`
	}
)

func NewConfiguration() (*EnvConfig, error) {
	conf := &EnvConfig{
		env: *envVar,
	}

	if err := loadConfigJson(*configJson, conf); err != nil {
		return nil, errors.Wrap(err, "failed to load config json")
	}

	return conf, nil
}

func (c *EnvConfig) GetEnv() string {
	return c.env
}

func loadConfigJson(filename string, config interface{}) error {
	if filename == "" {
		return errors.New("invalid config file path")
	}
	if config == nil {
		return errors.New("nil reference on config struct")
	}

	configData, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(configData, config)
}
