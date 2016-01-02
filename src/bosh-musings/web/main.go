package main

import (
	"bosh-musings/utils"
	"bosh-musings/web/handler"

	"log"
	"net/http"
	"os"
)

func main() {
	pidFile := os.Getenv("PIDFILE")
	err := utils.WritePIDFile(pidFile)
	if err != nil {
		log.Fatal("WritePIDFile:", err)
	}
	defer os.Remove(pidFile)

	http.HandleFunc("/", handler.ServeHTTP)
	err = http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}