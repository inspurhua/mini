package main

import (
	"fmt"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/router"
	"net/http"
)

func main() {
	r := router.NewRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", config.HTTPPort),
		Handler:        r,
		ReadTimeout:    config.ReadTimeout,
		WriteTimeout:   config.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
