package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"medods/config"
)

func getHandlers() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/home", homeHandler)
	r.HandleFunc("/user/{guidToken}", getUserTokens).Methods("GET")
	r.HandleFunc("/user/{guidToken}", updateUserTokens).Methods("POST")

	return r
}

func NewServer(conf *config.Config) *http.Server {
	//	Присваиваю  параметры серверу, которые были указаны в конфиге
	srv := &http.Server{
		Handler:      getHandlers(),
		Addr:         conf.Network.Address + conf.Network.Port,
		WriteTimeout: time.Duration(conf.Network.WriteTimeout) * time.Second,
		ReadTimeout:  time.Duration(conf.Network.ReadTimeout) * time.Second,
	}

	return srv
}
