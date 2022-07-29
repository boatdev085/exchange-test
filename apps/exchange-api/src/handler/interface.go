package handler

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	GetAssets(c echo.Context) error
	GetProfile(c echo.Context) error
	GetWallet(c echo.Context) error
	CreateOrder(c echo.Context) error
}
