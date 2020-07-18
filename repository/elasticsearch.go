package repository

import (
	"context"
	"encoding/json"
	"log"

	db "github.com/daniel-acaz/api-golang-elastic/config"
	model "github.com/daniel-acaz/api-golang-elastic/domain"
)

func FindAll() []model.Property {

	elasticsearch := db.GetConnection()

	res, err := elasticsearch.Search(
		elasticsearch.Search.WithContext(context.Background()),
		elasticsearch.Search.WithIndex("registries_index"),
		elasticsearch.Search.WithSize(600),
		elasticsearch.Search.WithTrackTotalHits(true),
		elasticsearch.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var response db.SearchResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	var properties []model.Property
	for _, hit := range response.Hits.Hits {
		registry := hit.Source
		properties = append(properties, registry)
	}

	return properties
}
