package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	db "scrap/database"
	line "scrap/lineNotify"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/gorm"
)

// Stock struct represents the data model for the stocks.
type Stock struct {
	ID             uint    `gorm:"primary_key"`
	StockSymbol    string  `gorm:"column:stock_symbol"`
	Price          float64 `gorm:"column:price"`
	PriceChange    float64 `gorm:"column:price_change"`
	PriceChangePct float64 `gorm:"column:price_change_pct"`
}

func main() {
	db := db.Connect()

	// Auto Migrate the Stock model
	db.AutoMigrate(&Stock{})

	// 股票代碼列表
	var stockSymbols = []string{"AAPL", "TSLA", "NVDA"}

	// 創建一個等待組，以確保所有 goroutine 都完成後才繼續
	var wg sync.WaitGroup

	// 創建一個 channel 來接收更新後的股票資料
	stockDataCh := make(chan Stock)

	// 啟動多個 goroutine 來處理不同的股票
	for _, symbol := range stockSymbols {
		wg.Add(1)
		go func(sym string) {
			defer wg.Done()

			// 獲取股票數據
			stockData, err := getStockData(db, sym)
			fmt.Println("傳入 channel 之前")
			if err != nil {
				log.Printf("Error getting stock data for %s:%v\n", sym, err)
				return
			}
			// 將股票資訊傳送到 channel
			stockDataCh <- stockData
		}(symbol)
	}

	// 等待所有 goroutine 完成並關閉 channel
	go func() {
		fmt.Println("等待 wg.wait")
		wg.Wait()
		fmt.Println("等待 結束")

		close(stockDataCh)
	}()

	// 接收 channel 數據並將其存到資料庫中
	fmt.Println("開始接收 channel 資訊")
	for stockData := range stockDataCh {
		if err := db.Save(&stockData).Error; err != nil {
			log.Println("Error save to database:", err)
		}
		line.Linenotify(stockData.StockSymbol, stockData.PriceChangePct)
		fmt.Printf("Save %s in the database\n", stockData.StockSymbol)
	}

	fmt.Println("Scraping and updating completed")

}

func getStockData(db *gorm.DB, symbol string) (Stock, error) {
	// 創建一個新的 collection
	c := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com"),
	)

	// 設置 http 請求之前的處理
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// 查詢資料庫裡是否有相同的股票，將資料庫已存在的股票資訊存給 existingStock 這個變數
	// 若過程發生錯誤，並且錯誤不是 record not found 的話，則返回錯誤
	// 反之如果錯誤是 record not found 就繼續執行程式
	var existingStock Stock
	if err := db.Where("stock_symbol = ?", symbol).Limit(1).Find(&existingStock).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return Stock{}, err
	}

	// 定義股票結構
	var stockData Stock

	// 設置傳回函數
	c.OnHTML("div[class='D(ib) Mend(20px)']", func(e *colly.HTMLElement) {
		// 抓取股票代號
		stockData.StockSymbol = symbol

		// 抓取價格相關訊息
		priceStr := e.ChildText("fin-streamer[data-test='qsp-price']")
		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			log.Println("Error parsing price:", err)
			return
		}
		stockData.Price = price

		// 抓取價格變化
		priceChangeStr := e.ChildText("fin-streamer[data-test='qsp-price-change'] span")
		priceChange, err := strconv.ParseFloat(priceChangeStr, 64)
		if err != nil {
			log.Println("Error parsing price change:", err)
			return
		}
		stockData.PriceChange = priceChange

		// 抓取價格百分比
		priceChangePctStr := e.ChildText("fin-streamer[data-field='regularMarketChangePercent'] span")
		priceChangePct, err := parseWithPercentSymbol(priceChangePctStr)
		if err != nil {
			log.Println("Error parsing price change percentage:", err)
			return
		}
		stockData.PriceChangePct = priceChangePct

	})

	// 訪問股票頁面
	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", symbol)
	err := c.Visit(url)
	if err != nil {
		return stockData, err
	}

	// 等待一段時間確保所有 http 請求都完成
	time.Sleep(2 * time.Second)

	// 檢查 existingStock.ID != 0，代表資料庫已存在相同資料，更新最新資料後返回
	// 如果 existingStock.ID == 0，代表沒有相同資料，直接返回原資料
	if existingStock.ID != 0 {
		fmt.Printf("Update %s data in the database.\n", existingStock.StockSymbol)
		existingStock.Price = stockData.Price
		existingStock.PriceChange = stockData.PriceChange
		existingStock.PriceChangePct = stockData.PriceChangePct
		// if err := db.Save(&existingStock).Error; err != nil {
		// 	return Stock{}, err
		// }
		return existingStock, nil
	} else {
		// if err := db.Create(&stockData).Error; err != nil {
		// 	return Stock{}, err
		// }
		return stockData, nil
	}
}

func parseWithPercentSymbol(value string) (float64, error) {
	// 使用正則表達式去除非數字字符，保留百分比符號
	re := regexp.MustCompile(`[^\d.-]+`)
	cleanedValue := re.ReplaceAllString(value, "")

	// 解析字符串為 float64
	return strconv.ParseFloat(cleanedValue, 64)
}
