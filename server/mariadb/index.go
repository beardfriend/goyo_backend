package mariadb

import (
	"fmt"
	"log"
	"sync"

	"goyo/libs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *gorm.DB
	onceDb   sync.Once
)

func GetInstance() *gorm.DB {
	onceDb.Do(func() {
		var err error

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			libs.ENV.Database.User, libs.ENV.Database.Pass, libs.ENV.Database.Host, libs.ENV.Database.Port, libs.ENV.Database.DB)

		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("Could not connect to database :%v", err)
		}
		instance = db
	})
	return instance
}
