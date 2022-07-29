package services

import (
	"exchange/apps/exchange-api/src/models"
	"exchange/apps/exchange-api/src/repositories"
	"time"

	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type sService struct {
	r  repositories.Repositories
	DB *gorm.DB
}

func NewService(r repositories.Repositories, d *gorm.DB) Services {
	return &sService{
		r:  r,
		DB: d,
	}
}

func (s *sService) GetAssets() ([]models.Assets, error) {
	assets, err := s.r.GetAssets()
	return assets, err

}

func (s *sService) GetProfile(userID int) (models.Profile, error) {
	profile, err := s.r.GetProfile(userID)
	return profile, err

}

func (s *sService) GetWallet(userID int) ([]models.Wallet, error) {
	wallet, err := s.r.GetWallet(userID)
	return wallet, err

}

func (s *sService) CreateOrder(dataOrder models.OrderParams) (int, error) {
	var dateCreated = null.TimeFrom(time.Now())
	var createOrder models.Order
	createOrder.Amount = dataOrder.Amount
	createOrder.DateCreated = dateCreated
	createOrder.AssetId = dataOrder.AssetId
	createOrder.UserId = dataOrder.UserId
	createOrder.OrderType = dataOrder.OrderType
	createOrder.PriceAction = dataOrder.PriceAction
	id, err := s.r.CreateOrder(s.DB, createOrder)

	return id, err
}
