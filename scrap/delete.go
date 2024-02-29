package scrap

import (
	"log"
	db "stockscrap/database"
	"stockscrap/database/models"

	"gorm.io/gorm"
)

func Deleter() {
	db := db.Connect()

	if err := deleteStockData(db); err != nil {
		log.Fatal("股票刪除失敗, 想不到連當韭菜都不夠格:", err)
	}

}

func deleteStockData(db *gorm.DB) error {
	// 抓取欲刪除的股票
	if err := db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Stock{}).Error; err != nil {
		return err
	}

	return nil
}
