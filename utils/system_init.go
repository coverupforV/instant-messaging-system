package utils

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config Mysql:", viper.Get("Mysql"))
	fmt.Println("config init success")
}

func InitMysql() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("Mysql.dsn")), &gorm.Config{})
	fmt.Println("mysql init success")
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println("user:", user)
}
