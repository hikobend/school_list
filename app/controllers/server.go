package controllers

import (
	"net/http"

	"example.com/school/config"
)

func StartmainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
