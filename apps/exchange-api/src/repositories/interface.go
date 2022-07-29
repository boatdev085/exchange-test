package repositories

import (
	"exchange/apps/exchange-api/src/models"

	"gorm.io/gorm"
)

type Repositories interface {
	GetAssets() ([]models.Assets, error)
	GetProfile(userID int) (models.Profile, error)
	GetWallet(userID int) ([]models.Wallet, error)
	CreateOrder(tx *gorm.DB, order models.Order) (int, error)
}
