package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"stone/transaction-challenge/api/middleware"
	_transactionRepo "stone/transaction-challenge/transaction/repository/api"
	_transactionRouter "stone/transaction-challenge/transaction/router"
	_transactionUcase "stone/transaction-challenge/transaction/usecase"
)

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := gin.Default()
	router.Use(middleware.Cors())

	transactionRepo := _transactionRepo.NewApiTransactionRepository(viper.GetString("pagarmeapi.baseUrl"))

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	transactionUcase := _transactionUcase.NewTransactionUsecase(transactionRepo, timeoutContext)
	_transactionRouter.NewTransactionRouter(router, transactionUcase)

	router.Run(viper.GetString("server.address"))
}
