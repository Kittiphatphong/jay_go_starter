package database

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func MySqlConvection() (*gorm.DB, error) {
	myDSN := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.database"),
	)
	fmt.Println("CONNECTING_TO_DB_MYSQL")
	//db := mysql.Open(myDSN)
	//sql, err := gorm.Open(db)
	openConnectionDB, err = gorm.Open(mysql.Open(myDSN), &gorm.Config{
		NowFunc: func() time.Time {
			tim, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(tim)
		},
	})
	if err != nil {
		log.Fatal("CONNECT_DATABASE_ERROR", err)
	}
	fmt.Println("DB_CONNECTED")
	return openConnectionDB, nil
}
