package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *Conf

type Conf struct {
	System *System `yaml:"system"`
	MySQL  *MySQL  `yaml:"mysql"`
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

func InitConfig() {
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
	fmt.Printf("%+v\n", Config.MySQL)
}
