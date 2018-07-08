package config

import (
	"github.com/spf13/viper"
	"strings"
	"github.com/fsnotify/fsnotify"
	"log"
)

type Config struct {
	Name string
}

func Init() error {
	c := Config{
		Name: "",
	}

	if err := c.initConfig(); err != nil {
		return err
	}
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		// 默认配置文件夹为conf, 文件名为config
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	// Define a prefix that ENVIRONMENT variables will use.
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
	})
}