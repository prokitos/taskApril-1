package services

import (
	"encoding/json"
	"io/ioutil"
	"module/internal/database"
	"module/internal/models"
	"net"
	"net/http"
	"net/url"
	"time"

	log "github.com/sirupsen/logrus"
)

// отправка номеров на внешний сервер по роуту /info
func sendRequestToGet(p_num string) models.Car {

	var strConnect string = ""
	database.GetExternalRoutes(&strConnect)

	baseURL, _ := url.Parse(strConnect + "/info")
	params := url.Values{}
	params.Add("regNum", p_num)
	baseURL.RawQuery = params.Encode()

	// для контекста
	client := http.Client{
		Timeout: 1 * time.Second,
	}

	// ctx, cncl := context.WithTimeout(context.Background(), time.Second*3)
	// defer cncl()
	// req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://google.com", nil)
	// resp, _ := http.DefaultClient.Do(req)

	resp, err := client.Get(baseURL.String())
	if err != nil {

		if err, ok := err.(net.Error); ok && err.Timeout() {
			log.Error("timeout request !!")
			var empty models.Car
			return empty
		}

		log.Debug("Error connecting to external api")
		log.Error("Error getting data from api")

		var empty models.Car
		return empty
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var carExemp models.Car
	json.Unmarshal(body, &carExemp)
	return carExemp
}
