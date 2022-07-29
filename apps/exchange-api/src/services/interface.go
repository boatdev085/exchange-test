package services

import (
	"exchange/apps/exchange-api/src/models"
)

type Services interface {
	GetAssets() ([]models.Assets, error)
	GetProfile(userID int) (models.Profile, error)
	GetWallet(userID int) ([]models.Wallet, error)
	CreateOrder(dataOrder models.OrderParams) (int, error)
}
