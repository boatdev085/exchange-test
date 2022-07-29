package main

import (
	"exchange/apps/exchange-api/src/connection"
	"exchange/apps/exchange-api/src/router"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(".env file not found")
	}

	exchangeDB := connection.NewDb(os.Getenv("DB_EXCHANGE"))

	r := router.New()
	router.GetRoutes(r, exchangeDB)

	r.Logger.Fatal(r.Start(":" + os.Getenv("PORT")))
}
