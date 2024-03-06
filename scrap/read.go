package scrap

import (
	"fmt"
	"log"
	db "stockscrap/database"
	"stockscrap/database/models"

	"gorm.io/gorm"
)

// 正常搜尋
func Reader(stockSymbols []string) {
	db := db.Connect()

	if err := readStockData(db, stockSymbols); err != nil {
		log.Fatal("查看股票失敗, 國防布搞屁:", err)
	}
}

// flag價格搜尋
func PriceReader(priceLimit int) {
	db := db.Connect()

	if err := priceStockData(db, priceLimit); err != nil {
		log.Fatal("價格甜不甜?", err)
	}
}

// 正常搜尋;顯示資料庫的股票
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

// flag價格搜尋;給 > N 價格以下的所有股票
func priceStockData(db *gorm.DB, priceLimit int) error {
	var stockDatas []models.Stock
	result := db.Where("price < ?", priceLimit).Find(&stockDatas)
	if result.Error != nil {
		log.Print("立 flag 錯了嗎?\n", result.Error)
		return result.Error
	}

	// 資料庫沒有找到設定值以下的股票
	if result.RowsAffected == 0 {
		fmt.Println("買什麼都漲的意思嗎? to the moon!")
		return nil
	}

	for _, stock := range stockDatas {
		fmt.Printf("股票代號: %s, 價格: %.1f, 漲跌價格: %.1f, 漲跌百分比: %.1f%%\n", stock.StockSymbol, stock.Price, stock.PriceChange, stock.PriceChangePct)
	}

	return nil
}
