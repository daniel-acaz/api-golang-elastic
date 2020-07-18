package config

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"

	model "github.com/daniel-acaz/api-golang-elastic/domain"
)

func GetConnection() *elasticsearch.Client {
	db, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	return db
}

type SearchResponse struct {
	Took int64
	Hits struct {
		Total struct {
			Value int64
		}
		Hits []*SearchHit
	}
}

type SearchHit struct {
	Score   float64 `json:"_score"`
	Index   string  `json:"_index"`
	Type    string  `json:"_type"`
	Version int64   `json:"_version,omitempty"`

	Source model.Property `json:"_source"`
}
