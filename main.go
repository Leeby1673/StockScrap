package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Stock struct represents the data model for the stocks.
type Stock struct {
	ID             uint   `gorm:"primary_key"`
	StockSymbol    string `gorm:"column:stock_symbol"`
	Price          float64
	PriceChange    float64 `gorm:"column:price_change"`
	PriceChangePct float64 `gorm:"column:price_change_pct"`
}

func main() {
	// Connect to MySQL database
	db, err := gorm.Open("mysql", "root:greed9527@tcp(localhost:3306)/stockscrap?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
			if err != nil {
				log.Printf("Error getting stock data for %s:%v\n", sym, err)
				return
			}
			// 將股票資訊傳送到 channel
			stockDataCh <- stockData
		}(symbol)
	}

	// 啟動一個 goroutine 來接收 channel 中的數據並將其存到資料庫中
	go func() {
		for stockData := range stockDataCh {
			if err := db.Save(&stockData).Error; err != nil {
				log.Println("Error save to database:", err)
			}
		}
		// // 檢查資料庫中是否已存在相同的股票
		// var existingStock Stock
		// if db.Where("stock_symbol = ?", dataSymbol).First(&existingStock).RecordNotFound() {
		// 	// 如果不存在，則儲存到資料庫
		// 	stock := Stock{
		// 		StockSymbol:    dataSymbol,
		// 		Price:          price,
		// 		PriceChange:    priceChange,
		// 		PriceChangePct: priceChangePct,
		// 	}

		// 	// Save to database
		// 	if err := db.Save(&stock).Error; err != nil {
		// 		log.Println("Error saving to database:", err)
		// 	}
		// } else {
		// 	fmt.Printf("Stock %s already exists in the database.\n", dataSymbol)
		// }
	}()

	// 等待所有 goroutine 完成
	wg.Wait()

	// 關閉 channel
	close(stockDataCh)

	fmt.Println("Scraping and updating completed")

}

func getStockData(db *gorm.DB, symbol string) (Stock, error) {
	// 創建一個新的 collection
	c := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com"),
	)

	// 查詢資料庫裡是否有相同的股票，將資料庫已存在的股票資訊存給 existingStock 這個變數
	// 若過程發生錯誤，並且錯誤不是 record not found 的話，則返回錯誤
	var existingStock Stock
	if err := db.Where("stock_symbol = ?", symbol).First(&existingStock).Error; err != nil && !gorm.IsRecordNotFoundError(err) {
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

	// 檢查 existingStock.ID != 0，代表資料庫已存在相同資料，並更新現有資訊
	// 如果 existingStock.ID == 0，代表沒有相同資料，就創建一個
	if existingStock.ID != 0 {
		existingStock.Price = stockData.Price
		existingStock.PriceChange = stockData.PriceChange
		existingStock.PriceChangePct = stockData.PriceChangePct
		if err := db.Save(&existingStock).Error; err != nil {
			return Stock{}, err
		}
		return existingStock, nil
	} else {
		if err := db.Create(&stockData).Error; err != nil {
			return Stock{}, err
		}
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
