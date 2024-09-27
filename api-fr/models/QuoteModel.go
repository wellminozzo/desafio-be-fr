package models

import (
	"gorm.io/gorm"
)

type Metric struct {
	CompaniesMetrics CompanyMetric `json:"companies_metrics"`
	CheaperQuote     Quote         `json:"cheaper_quote"`
	ExpensiveQuote   Quote         `json:"expensive_quote"`
}

type CompanyMetric struct {
	Name        string  `json:"name"`
	TotalPrice  float64 `json:"total_price"`
	AvgPrice    float64 `json:"avg_price"`
	QuotesCount uint8   `json:"quotes_count"`
}
type Quote struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `json:"name"`
	Price    string `json:"price"`
	Service  string `json:"service"`
	Deadline string `json:"deadline"`
}

func GetQuoteStats(db *gorm.DB) ([]CompanyMetric, error) {
	var stats []CompanyMetric

	// Executa a query de agregação
	err := db.Table("quotes").
		Select("name, COUNT(*) as quotes_count, SUM(price) as soma_total_prices, AVG(price) as media_prices").
		Group("name").
		Scan(&stats).Error

	if err != nil {
		return nil, err
	}

	return stats, nil
}

func GetCheaperQuote(db *gorm.DB) (Quote, error) {
	var cheaperQuote Quote
	err := db.Table("quotes").
		Select("name, price, service, deadline").
		Order("price ASC").
		Limit(1).
		Scan(&cheaperQuote).Error
	if err != nil {
		return Quote{}, err
	}
	return cheaperQuote, nil
}

func GetExpensiveQuote(db *gorm.DB) (Quote, error) {
	var expensiveQuote Quote
	err := db.Table("quotes").
		Select("name, price, service, deadline").
		Order("price DESC").
		Limit(1).
		Scan(&expensiveQuote).Error
	if err != nil {
		return Quote{}, err
	}
	return expensiveQuote, nil
}
