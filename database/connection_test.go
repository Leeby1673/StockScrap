package db

import (
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestConnectWithGorm(t *testing.T) {
	// 連接到 MySQL 資料庫
	dsn := "root:greed9527@tcp(localhost:3306)/stockscrap?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("無法連接到資料庫：%v", err)
	}

	// 檢查資料庫連接狀態
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("資料庫連接失敗：%v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("資料庫 Ping 失敗：%v", err)
	}

	log.Println("資料庫連接成功！")
}
