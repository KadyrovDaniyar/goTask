package usecase

import (
	"course/domain"
	"course/repository"
	"log"
	"time"

	"github.com/spf13/viper"
)

func UpdatePrices(coins *[]domain.Coin) {
	// бесконечный цикл обновления курсов
	for {
		newCoins, err := repository.GetCoins()
		if err != nil {
			log.Printf("Error updating prices: %s\n", err)
		} else {
			*coins = newCoins
			log.Println("Prices updated")
		}
		time.Sleep(time.Duration(viper.GetInt("APP.COIN.TIMEOUT.MIN")) * time.Minute)
	}
}
