package domain

import "time"

type Coin struct {
	ID                string    `json:"id"`
	Symbol            string    `json:"symbol"`
	Name              string    `json:"name"`
	Image             string    `json:"image"`
	CurrentPrice      float64   `json:"current_price"`
	MarketCap         float64   `json:"market_cap"`
	MarketCapRank     int       `json:"market_cap_rank"`
	TotalVolume       float64   `json:"total_volume"`
	High24            float64   `json:"high_24h"`
	Low24             float64   `json:"low_24h"`
	PriceChange24     float64   `json:"price_change_24h"`
	PriceChangePct24h float64   `json:"price_change_percentage_24h"`
	LastUpdated       time.Time `json:"last_updated"`
}
