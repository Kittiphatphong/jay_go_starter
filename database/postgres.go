package database

import (
	"fmt"
	"go_starter/config"
	"go_starter/logs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

type SqlLogger struct {
	logger.Interface
}

var openConnectionDB *gorm.DB
var err error

func PostgresConnection() (*gorm.DB, error) {
	myDSN := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Bangkok",
		config.Env("postgres.host"),
		config.Env("postgres.user"),
		config.Env("postgres.password"),
		config.Env("postgres.database"),
		config.Env("postgres.port"),
	)

	fmt.Println("CONNECTING_TO_POSTGRES_DB")
	openConnectionDB, err = gorm.Open(postgres.Open(myDSN), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			ti, _ := time.LoadLocation("Asia/Bangkok")
			return time.Now().In(ti)
		},
	})
	//DryRun: false,
	if err != nil {
		logs.Error(err)
		log.Fatal("ERROR_PING_POSTGRES", err)
		return nil, err
	}
	fmt.Println("POSTGRES_CONNECTED")
	return openConnectionDB, nil
}
