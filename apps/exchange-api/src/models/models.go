package models

import (
	"github.com/shopspring/decimal"
	"gopkg.in/guregu/null.v4"
)

type Assets struct {
	ID          int             `json:"id"`
	DateCreated null.Time       `json:"date_created"`
	DateUpdated null.Time       `json:"date_updated"`
	LastPrice   decimal.Decimal `json:"last_price"`
	Logo        string          `json:"logo"`
	Name        string          `json:"name"`
	Symbol      string          `json:"symbol"`
}

type Profile struct {
	ID          int       `json:"id"`
	DateCreated null.Time `json:"date_created"`
	DateUpdated null.Time `json:"date_updated"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
}

type Wallet struct {
	DateCreated null.Time       `json:"date_created"`
	DateUpdated null.Time       `json:"date_updated"`
	Amount      decimal.Decimal `json:"amount"`
	AssetId     int             `json:"asset_id"`
	Symbol      string          `json:"symbol"`
}

type OrderParams struct {
	AssetId     int     `json:"asset_id"`
	UserId      int     `json:"user_id"`
	OrderType   string  `json:"order_type"`
	PriceAction float64 `json:"price_action"`
	Amount      float64 `json:"amount"`
}

type Order struct {
	DateCreated null.Time `json:"date_created"`
	DateUpdated null.Time `json:"date_updated"`
	AssetId     int       `json:"asset_id"`
	UserId      int       `json:"user_id"`
	OrderType   string    `json:"order_type"`
	PriceAction float64   `json:"price_action"`
	Amount      float64   `json:"amount"`
}
