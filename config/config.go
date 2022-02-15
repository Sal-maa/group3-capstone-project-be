package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port       int    `json:"port" yaml:"port"`
	JWT_secret string `json:"secret" yaml:"secret"`
	Database   struct {
		Driver   string `json:"driver" yaml:"driver"`
		Host     string `json:"host" yaml:"host"`
		Port     int    `json:"port" yaml:"port"`
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		Name     string `json:"name" yaml:"name"`
	}
	AWS struct {
		AccessKeyID string `json:"accesskeyid" yaml:"accesskeyid"`
		SecretKey   string `json:"secretkey" yaml:"secretkey"`
		Region      string `json:"region" yaml:"region"`
		Bucket      string `json:"bucket" yaml:"bucket"`
	}
}

var lock *sync.Mutex
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock = &sync.Mutex{}
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	defaultConfig := AppConfig{}

	// set app default configuration in aws ec2
	defaultConfig.Port, _ = strconv.Atoi(os.Getenv("PORT"))
	defaultConfig.JWT_secret = os.Getenv("JWT_SECRET")
	defaultConfig.Database.Driver = os.Getenv("DB_DRIVER")
	defaultConfig.Database.Host = os.Getenv("DB_HOST")
	defaultConfig.Database.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	defaultConfig.Database.Username = os.Getenv("DB_USERNAME")
	defaultConfig.Database.Password = os.Getenv("DB_PASSWORD")
	defaultConfig.Database.Name = os.Getenv("DB_NAME")
	defaultConfig.AWS.AccessKeyID = os.Getenv("AWS_ACCESSKEYID")
	defaultConfig.AWS.SecretKey = os.Getenv("AWS_SECRETKEY")
	defaultConfig.AWS.Region = os.Getenv("AWS_REGION")
	defaultConfig.AWS.Bucket = os.Getenv("AWS_BUCKET")

	// viper.SetConfigType("json")
	// viper.SetConfigName("config")
	// viper.AddConfigPath("./config")

	// // read app custom configuration for running in local machine
	// if err := viper.ReadInConfig(); err != nil {
	// 	log.Println(err)
	// 	return &defaultConfig
	// }

	// finalConfig := AppConfig{}

	// if err := viper.Unmarshal(&finalConfig); err != nil {
	// 	log.Println(err)
	// 	return &defaultConfig
	// }

	// return &finalConfig

	return &defaultConfig
}
