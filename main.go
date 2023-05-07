package main

import (
	"course/domain"
	"course/usecase"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"
)

var coins []domain.Coin //global variable

func main() {

	// Config
	viper.SetConfigName("config.yml")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error loading config: %s\n", err)
	}

	go usecase.UpdatePrices(&coins)

	http.HandleFunc("/coins", getCoins)
	http.HandleFunc("/coin", getCoinByID) //coin?id=bitcoin
	if err := http.ListenAndServe(":"+viper.GetString("APP.PORT"), nil); err != nil {
		fmt.Println("HTTP server error:", err)
	}
}

func getCoins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coins)
}

func getCoinByID(w http.ResponseWriter, r *http.Request) {
	coinName := r.URL.Query().Get("id")
	if coinName == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	for _, coin := range coins {
		if coin.Name == coinName || coin.ID == coinName || coin.ID == "bitcoin" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(coin)
			return
		}
	}

	http.NotFound(w, r)
}
