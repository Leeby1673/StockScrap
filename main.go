package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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

	// Create a new collector
	c := colly.NewCollector(
		colly.AllowedDomains("finance.yahoo.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		// You can set request headers or other properties here if needed
		fmt.Println("Visiting", r.URL)
	})

	// Define the URL to scrape
	url := "https://finance.yahoo.com/quote/%5EGSPC?p=^GSPC"

	// Setup the collector to visit the URL and extract stock information
	c.OnHTML("div[class='D(ib) Mend(20px)']", func(e *colly.HTMLElement) {
		stockSymbol := e.ChildText("h1")
		priceStr := e.ChildText("span[data-reactid='50']")
		priceChangeStr := e.ChildText("span[data-reactid='51']")
		priceChangePctStr := e.ChildText("span[data-reactid='52']")

		price, err := strconv.ParseFloat(strings.Replace(priceStr, ",", "", -1), 64)
		if err != nil {
			log.Println("Error parsing price:", err)
			return
		}

		priceChange, err := strconv.ParseFloat(strings.Replace(priceChangeStr, ",", "", -1), 64)
		if err != nil {
			log.Println("Error parsing price change:", err)
			return
		}

		priceChangePct, err := strconv.ParseFloat(strings.Replace(priceChangePctStr, "%", "", -1), 64)
		if err != nil {
			log.Println("Error parsing price change percentage:", err)
			return
		}

		// Save the data to the database
		stock := Stock{
			StockSymbol:    stockSymbol,
			Price:          price,
			PriceChange:    priceChange,
			PriceChangePct: priceChangePct,
		}

		// Save to database
		if err := db.Save(&stock).Error; err != nil {
			log.Println("Error saving to database:", err)
		}
	})

	// Start scraping
	err = c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scraping completed!")
}
