package repository

import (
	"course/domain"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetCoins() (coins []domain.Coin, err error) {

	// выполняем запрос к API
	resp, err := http.Get(viper.GetString("APP.COIN.URL"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// читаем ответ в виде []byte
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// декодируем JSON
	err = json.Unmarshal(body, &coins)
	if err != nil {
		return nil, err
	}

	return coins, nil
}
