package db

import (
	"log"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// var DB *gorm.DB

func Connect() *gorm.DB {
	// Connect to MySQL database
	dsn := "root:greed9527@tcp(localhost:3306)/stockscrap?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Hour) // 每條連線的存活時間
	sqlDB.SetMaxOpenConns(8)            // 最大連線數
	sqlDB.SetMaxIdleConns(6)            // 最大閒置連線數

	return db
}
