package scrap

import (
	"fmt"
	"log"
	db "stockscrap/database"
	"stockscrap/database/models"

	"gorm.io/gorm"
)

func Deleter(stockSymbols []string) {
	db := db.Connect()

	if err := deleteStockData(db, stockSymbols); err != nil {
		log.Fatal("股票刪除失敗, 想不到連當韭菜都不夠格:", err)
	}

}

func deleteStockData(db *gorm.DB, stockSymbols []string) error {
	var stockDatas []models.Stock
	// 抓取欲刪除的股票
	for _, symbol := range stockSymbols {
		if err := db.Where("stock_symbol = ?", symbol).Delete(&stockDatas).Error; err != nil {
			return err
		}
		fmt.Printf("刪除 %s, 解套了嗎?\n", symbol)
	}

	return nil
}
