package config

import (
	"os"
	"sync"

	"github.com/spf13/viper"
)

var (
	Config   *Conf
	initOnce sync.Once
)

type Conf struct {
	System    *System    `yaml:"system"`
	MySQL     *MySQL     `yaml:"mysql"`
	Email     *Email     `yaml:"email"`
	LogConfig *LogConfig `yaml:"logConfig"`
}

type System struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"userName"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

type Email struct {
	SmtpHost  string `yaml:"smtpHost"`
	SmtpEmail string `yaml:"smtpEmail"`
	SmtpPass  string `yaml:"smtpPass"`
}

type LogConfig struct {
	Level    string `yaml:"level"`
	FilePath string `yaml:"filePath"`
}

// GetConfig 获取配置，如果配置未初始化则自动初始化
func GetConfig() *Conf {
	if Config == nil {
		InitConfig()
	}
	return Config
}

// InitConfig 初始化配置，使用sync.Once确保只初始化一次
func InitConfig() {
	initOnce.Do(func() {
		workDir, _ := os.Getwd()
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(workDir + "/config/locales")
		viper.AddConfigPath(workDir)

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
		err = viper.Unmarshal(&Config)
		if err != nil {
			panic(err)
		}
	})
}

// ResetConfig 重置配置，主要用于测试
func ResetConfig() {
	Config = nil
}
