package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v8/esapi"

	db "github.com/daniel-acaz/api-golang-elastic/config"
	model "github.com/daniel-acaz/api-golang-elastic/domain"
)

func FindAll() []model.Property {

	elasticsearch := db.GetConnection()

	res, err := elasticsearch.Search(
		elasticsearch.Search.WithContext(context.Background()),
		elasticsearch.Search.WithIndex("properties_index"),
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
		property := hit.Source
		properties = append(properties, property)
	}

	return properties
}

func FindById(id int) model.Property {

	elasticsearch := db.GetConnection()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := elasticsearch.Search(
		elasticsearch.Search.WithContext(context.Background()),
		elasticsearch.Search.WithIndex("properties_index"),
		elasticsearch.Search.WithBody(&buf),
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
		property := hit.Source
		properties = append(properties, property)
	}

	return properties[0]
}

func Save(property model.Property) model.Property {

	elasticsearch := db.GetConnection()

	body, err := json.Marshal(property)
	if err != nil {
		log.Fatalf("Error parsing the registry: %s", err)
	}

	req := esapi.IndexRequest{
		Index:      "properties_index",
		DocumentID: strconv.Itoa(property.ID),
		Body:       strings.NewReader(string(body)),
		Refresh:    "true",
	}

	res, err := req.Do(context.Background(), elasticsearch)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document ID=%v", res.Status(), property.ID)
	} else {
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}

	return property

}

func GetMaxId() int {

	elasticsearch := db.GetConnection()

	res, err := elasticsearch.Search(
		elasticsearch.Search.WithContext(context.Background()),
		elasticsearch.Search.WithBody(strings.NewReader(`{
				"sort": [
				  {
					"id": {
					  "order": "desc"
					}
				  }
				],
				"size": 1
			  }`)),
		elasticsearch.Search.WithIndex("properties_index"),
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
		property := hit.Source
		properties = append(properties, property)
	}

	return properties[0].ID

}

func Delete(id int) {

	elasticsearch := db.GetConnection()

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"id": id,
			},
		},
	}

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	index := []string{"properties_index"}

	req := esapi.DeleteByQueryRequest{Index: index, Body: &buf}
	res, err := req.Do(context.Background(), elasticsearch)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}

	if res.IsError() {
		log.Fatal(res.String())
	}
}
