package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// запуск сервера
func MainServer() {

	router := routers()

	srv := &http.Server{
		Handler:      router,
		Addr:         ":8888",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}

// все пути
func routers() *mux.Router {
	router := mux.NewRouter()

	// потом можно поменять все пути на cars, если будет несколько таблиц
	router.HandleFunc("/delete", deleteRequest).Methods(http.MethodDelete)
	router.HandleFunc("/insert", insertGetRequest).Methods(http.MethodPost)
	router.HandleFunc("/update", updateGetRequest).Methods(http.MethodPut)
	router.HandleFunc("/show", showsSpecGetRequest).Methods(http.MethodGet)

	return router
}
