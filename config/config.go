package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

const (
	defaultConfigFilename = "config.yaml"
	configEnvKey          = "JNN_JRS"
)

const (
	CTX_KEY_USER_ID = "useId"
)

var Conf Config

type Config struct {
	System systemConf `mapstructure:"system" json:"system" yaml:"system"`
	Zap    zapConf    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql  mysqlConf  `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Ldap   ldap       `mapstructure:"ldap" json:"ldap" yaml:"ldap"`
	Jwt    jwt        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Excel  excel      `mapstructure:"excel" json:"excel" yaml:"excel"`
}

type systemConf struct {
	Addr string `mapstructure:"addr" json:"addr" yaml:"addr"`
}

type ldap struct {
	URL string `mapstructure:"url" json:"url" yaml:"url"`
}

type jwt struct {
	Duration int    `mapstructure:"duration" json:"duration" yaml:"duration"`
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`
}

type excel struct {
	MasterResPlanPrefix string `mapstructure:"master-res-plan-prefix" json:"masterResPlanPrefix" yaml:"master-res-plan-prefix"`
}

func init() {
	v := viper.New()
	v.SetConfigFile(getConfigFilePath())
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	unmarshalConfig(v)
	// watching and updating Conf without application restart
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) { unmarshalConfig(v) })
}

func unmarshalConfig(v *viper.Viper) {
	if err := v.Unmarshal(&Conf); err != nil {
		panic(fmt.Errorf("fail to unmarshal conf file: %s", err))
	}
	// TODO: disconnect then connect
	initLogger()
	initMysql()
}

// getConfigFilePath get configuration file path
// priority: command line >> environment variable >> default value
func getConfigFilePath() (config string) {
	// from command line
	flag.StringVar(&config, "c", "", "input config file path")
	flag.Parse()
	if config != "" {
		fmt.Println("Config file passing from command line:", config)
		return
	}
	// from env var
	if env := os.Getenv(configEnvKey); env != "" {
		config = env
		fmt.Println("Config file passing from environment variable:", config)
		return
	}
	// from default value
	config = defaultConfigFilename
	fmt.Println("Default config file:", config)
	return
}
