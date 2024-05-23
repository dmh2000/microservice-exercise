package main

import (
	"log"
	"net/http"

	"sqirvy.xyz/microservice/movie/controller/movie"
	metadatagateway "sqirvy.xyz/microservice/movie/gateway/metadata/http"
	ratinggateway "sqirvy.xyz/microservice/movie/gateway/rating/http"
	httphandler "sqirvy.xyz/microservice/movie/handler/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
