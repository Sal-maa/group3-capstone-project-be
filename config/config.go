package config

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type AppConfig struct {
	Port       int    `json:"app_port"`
	JWT_secret string `json:"jwt_secret"`
	Database   struct {
		Driver   string `json:"db_driver"`
		Host     string `json:"db_host"`
		Port     int    `json:"db_port"`
		Username string `json:"db_username"`
		Password string `json:"db_password"`
		Name     string `json:"db_name"`
	}
	AWS struct {
		AccessKeyID string `json:"aws_accesskeyid"`
		SecretKey   string `json:"aws_secretkey"`
		Region      string `json:"aws_region"`
		Bucket      string `json:"aws_bucket"`
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

	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.AddConfigPath("./config/")

	// read app custom configuration for running in local machine
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
		return &defaultConfig
	}

	finalConfig := AppConfig{}

	if err := viper.Unmarshal(&finalConfig); err != nil {
		log.Println(err)
		return &defaultConfig
	}

	return &finalConfig
}
