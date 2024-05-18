package main

import (
	"log"
	"net/http"

	rating "sqirvy.xyz/microservice/rating/controller"
	httphandler "sqirvy.xyz/microservice/rating/handler"
	"sqirvy.xyz/microservice/rating/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
