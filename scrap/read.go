package scrap

import (
	"fmt"
	"log"
	db "stockscrap/database"
	"stockscrap/database/models"

	"gorm.io/gorm"
)

func Reader() {
	db := db.Connect()

	if err := readStockData(db); err != nil {
		log.Fatal("查看股票失敗, 國防布搞屁:", err)
	}
}

// 顯示資料庫的股票
func readStockData(db *gorm.DB) error {
	var stockDatas []models.Stock
	if err := db.Find(&stockDatas).Error; err != nil {
		return err
	}

	for _, stock := range stockDatas {
		fmt.Printf("股票代號: %s, 價格: %.1f, 漲跌價格: %.1f, 漲跌百分比: %.1f%%\n", stock.StockSymbol, stock.Price, stock.PriceChange, stock.PriceChangePct)
	}

	return nil
}
