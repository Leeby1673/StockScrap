package models

// Stock struct represents the data model for the stocks.
type Stock struct {
	ID             uint    `gorm:"primary_key"`
	StockSymbol    string  `gorm:"column:stock_symbol"`
	Price          float64 `gorm:"column:price"`
	PriceChange    float64 `gorm:"column:price_change"`
	PriceChangePct float64 `gorm:"column:price_change_pct"`
}
