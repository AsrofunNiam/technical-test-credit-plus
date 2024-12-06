package repository_test

import (
	"log"

	"github.com/AsrofunNiam/technical-test-credit-plus/app"
	c "github.com/AsrofunNiam/technical-test-credit-plus/configuration"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	configuration, err := loadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	db := app.ConnectDatabase(configuration.User, configuration.Host, configuration.Password, configuration.PortDB, configuration.Db)
	return db
}

func loadConfig() (config c.Configuration, err error) {
	viper.SetConfigFile("../configuration/.env")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
