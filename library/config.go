package library

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

// 初始化
func (c *Config) initialization() error {
	if c.Name != "" {
		// 如果指定了配置文件，则解析指定的配置文件
		viper.SetConfigFile("conf/" + c.Name + ".yaml")
	} else {
		// 如果没有指定配置文件，则解析默认的配置文件
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// 监听配置文件是否改变,用于热更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件改变: %s\n", e.Name)
	})
}

//获取配置 key
func (c *Config) Get(key string) interface{} {
	name := viper.Get(key)
	return name
}

// 调用方法
func NewConfig(cfg string) (*Config) {
	c := &Config{
		Name: cfg,
	}
	// 初始化配置文件
	err := c.initialization()
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// 热更新
	c.watchConfig()
	return c
}
