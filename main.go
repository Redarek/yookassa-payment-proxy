package main

import (
	"fmt"
	"os"
	"strconv"
	"yookassa-payment-proxy/internal/server"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	app := server.New()

	app.RegisterFiberRoutes()
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
