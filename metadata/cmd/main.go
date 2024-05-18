package main

import (
	"log"
	"net/http"

	"sqirvy.xyz/microservice/metadata/controller"
	httphandler "sqirvy.xyz/microservice/metadata/handler"
	memory "sqirvy.xyz/microservice/metadata/repository/memory"
)

func main() {
	log.Println("Starting the movie metadata service")
	repo := memory.New()
	ctrl := controller.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
