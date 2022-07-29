package router

import (
	"exchange/apps/exchange-api/src/handler"
	"exchange/apps/exchange-api/src/repositories"
	"exchange/apps/exchange-api/src/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetRoutes(e *echo.Echo, DBgorm *gorm.DB) {
	repo := repositories.NewRepo(DBgorm)

	ser := services.NewService(repo, DBgorm)

	h := handler.NewHandler(ser)

	gAPI := e.Group("/")
	createRouter(gAPI, h)
}

func createRouter(g *echo.Group, h handler.Handlers) {
	g.GET("assets", h.GetAssets)
	g.GET("profile/:userID", h.GetProfile)
	g.GET("wallet/:userID", h.GetWallet)
	g.POST("order", h.CreateOrder)

}
