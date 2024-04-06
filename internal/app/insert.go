package app

import (
	"encoding/json"
	"fmt"
	"io"
	"module/internal/models"
	"net/http"
)

func insertGetRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reqBody, _ := io.ReadAll(r.Body)
	var nomera models.CarNumber
	json.Unmarshal(reqBody, &nomera)

	for i := 0; i < 40; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	//services.RestCreateData(w, &nomera)
}
