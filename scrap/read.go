package scrap

import (
	"fmt"
	"log"
	db "stockscrap/database"
	"stockscrap/database/models"

	"gorm.io/gorm"
)

func Reader(stockSymbols []string) {
	db := db.Connect()

	if err := readStockData(db, stockSymbols); err != nil {
		log.Fatal("查看股票失敗, 國防布搞屁:", err)
	}
}

// 顯示資料庫的股票
func readStockData(db *gorm.DB, stockSymbols []string) error {
	var stockDatas []models.Stock
	// 若使用沒有提供參數，給資料庫全部的股票資訊
	// 有提供指定參數的話，就給參數指定的股票資訊
	if len(stockSymbols) == 0 {
		if err := db.Find(&stockDatas).Error; err != nil {
			return err
		}

		for _, stock := range stockDatas {
			fmt.Printf("股票代號: %s, 價格: %.1f, 漲跌價格: %.1f, 漲跌百分比: %.1f%%\n", stock.StockSymbol, stock.Price, stock.PriceChange, stock.PriceChangePct)
		}
	} else {
		for _, symbol := range stockSymbols {
			if err := db.Where("stock_symbol = ?", symbol).First(&stockDatas).Error; err != nil {
				log.Printf("資料庫內找不到 %s\n", symbol)
				return err
			}
			for _, stock := range stockDatas {
				fmt.Printf("股票代號: %s, 價格: %.1f, 漲跌價格: %.1f, 漲跌百分比: %.1f%%\n", stock.StockSymbol, stock.Price, stock.PriceChange, stock.PriceChangePct)
			}
		}

	}

	return nil
}
