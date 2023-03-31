package main

import (
	"github.com/favetelinguis/henke-larsson-ab-monorepo/handlers/rest"
	"log"
	"net/http"
)

func main() {
	addr := ":8080"
	mux := http.NewServeMux()
	mux.HandleFunc("/hellooo", rest.TranslateHandler)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
