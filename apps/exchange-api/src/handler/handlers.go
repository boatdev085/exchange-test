package handler

import (
	"exchange/apps/exchange-api/src/models"
	"exchange/apps/exchange-api/src/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type hanlers struct {
	s services.Services
}

func NewHandler(s services.Services) Handlers {
	return &hanlers{
		s: s,
	}
}

func (h *hanlers) GetAssets(c echo.Context) error {
	assets, err := h.s.GetAssets()
	if err != nil {
		return c.JSON(400, models.FailResponse("Invalid assets."))
	}
	return c.JSON(200, models.SuccessResponse("Get data success.", assets))
}

func (h *hanlers) GetProfile(c echo.Context) error {
	getUserID := c.Param("userID")
	userID, err := strconv.Atoi(getUserID)
	if err != nil {
		return c.JSON(400, models.FailResponse("Invalid user id."))
	}
	profile, err := h.s.GetProfile(userID)
	if err != nil {
		return c.JSON(400, models.FailResponse("Invalid profile."))
	}
	return c.JSON(200, models.SuccessResponse("Get data success.", profile))
}

func (h *hanlers) GetWallet(c echo.Context) error {
	getUserID := c.Param("userID")
	userID, err := strconv.Atoi(getUserID)
	if err != nil {
		return c.JSON(400, models.FailResponse("Invalid user id."))
	}
	wallet, err := h.s.GetWallet(userID)
	if err != nil {
		return c.JSON(400, models.FailResponse("Invalid wallet."))
	}
	return c.JSON(200, models.SuccessResponse("Get data success.", wallet))
}

func (h *hanlers) CreateOrder(c echo.Context) error {
	var err error
	var dataCreateOrder models.OrderParams
	if err = c.Bind(&dataCreateOrder); err != nil {
		return c.JSON(400, models.FailResponse("Invalid data format."))
	}

	if dataCreateOrder.OrderType == "" || dataCreateOrder.AssetId == 0 || dataCreateOrder.PriceAction == 0 || dataCreateOrder.UserId == 0 {
		return c.JSON(400, models.FailResponse("Invalid data format."))
	}
	id, err := h.s.CreateOrder(dataCreateOrder)
	if err != nil {
		return c.JSON(400, models.FailResponse("Create not working."))

	}
	return c.JSON(200, models.SuccessResponse("Get data success.", id))
}
