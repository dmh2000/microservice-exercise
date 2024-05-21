package model

import "sqirvy.xyz/microservice/metadata/pkg/model"

// MovieDetails includes movie metadata and its aggregated
// rating.
type MovieDetails struct {
	Rating   *float64       `json:"rating"`
	Metadata model.Metadata `json:"metadata"`
}
