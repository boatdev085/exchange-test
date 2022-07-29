package repositories

import (
	"exchange/apps/exchange-api/src/models"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type rRepo struct {
	DB *gorm.DB
}

func NewRepo(DB *gorm.DB) Repositories {
	return &rRepo{
		DB: DB,
	}
}

func (r *rRepo) GetAssets() ([]models.Assets, error) {

	var assets []models.Assets
	qString := `select id,symbol,date_created,date_updated,logo,name,last_price from assets`
	err := r.DB.Raw(qString).Scan(&assets).Error
	if err != nil {
		return assets, err
	}
	return assets, nil
}

func (r *rRepo) GetProfile(userID int) (models.Profile, error) {
	var profile models.Profile
	rows, err := r.DB.Table("users").Select(
		"id", "date_created", "first_name", "last_name").Where("id = ?", userID).Rows()
	if err != nil {
		return profile, err
	}
	for rows.Next() {
		err := rows.Scan(
			&profile.ID, &profile.DateCreated, &profile.FirstName,
			&profile.LastName,
		)
		if err != nil {
			return profile, err
		}
	}
	return profile, nil
}

func (r *rRepo) GetWallet(userID int) ([]models.Wallet, error) {
	var wallet []models.Wallet
	qString := `SELECT wallets.date_created,wallets.amount,assets.id as asset_id,assets.symbol as symbol FROM "wallets" inner Join assets on assets.id = wallets.asset_id WHERE user_id = ?`
	err := r.DB.Raw(qString, userID).Scan(&wallet).Error
	if err != nil {
		return wallet, err
	}
	return wallet, nil
}

func (r *rRepo) CreateOrder(tx *gorm.DB, order models.Order) (int, error) {
	var insertID int
	var col = []string{
		"asset_id",
		"date_created",
		"user_id", "order_type",
		"amount", "price_action",
	}
	qString := fmt.Sprintf(
		`INSERT INTO orders (%s) VALUES (%s) RETURNING id`,
		strings.Join(col, ","),
		setStringValueQuery(len(col)),
	)
	err := tx.Raw(qString,
		order.AssetId, order.DateCreated,
		order.UserId, order.OrderType,
		order.Amount, order.PriceAction,
	).Scan(&insertID).Error
	UpdateBalance(tx, order)
	return insertID, err
}

func UpdateBalance(tx *gorm.DB, dataOrder models.Order) error {

	if dataOrder.OrderType == "BUY" {
		qString := `UPDATE wallets SET 
			amount= amount + ?
			WHERE user_id = ? and asset_id = ?`

		tx.Exec(qString, dataOrder.Amount,
			dataOrder.UserId, dataOrder.AssetId,
		)
		qStringTHB := `UPDATE wallets SET 
		amount= amount - ?
		WHERE user_id = ? and asset_id = 3`

		tx.Exec(qStringTHB, dataOrder.PriceAction,
			dataOrder.UserId,
		)
	}
	if dataOrder.OrderType == "SELL" {
		qString := `UPDATE wallets SET 
			amount= amount - ?
			WHERE user_id = ? and asset_id = ?`

		tx.Exec(qString, dataOrder.Amount,
			dataOrder.UserId, dataOrder.AssetId,
		)
		qStringTHB := `UPDATE wallets SET 
		amount= amount + ?
		WHERE user_id = ? and asset_id = 3`

		tx.Exec(qStringTHB, dataOrder.PriceAction,
			dataOrder.UserId,
		)
	}
	return nil
}

func setStringValueQuery(lenValue int) string {
	var valueSet string
	for i := 1; i <= lenValue; i++ {
		valueSet += "?"
		if i < lenValue {
			valueSet += ","
		}
	}
	return valueSet
}
