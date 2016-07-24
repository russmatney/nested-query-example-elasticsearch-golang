package main

import (
	"fmt"

	"github.com/russmatney/nested-query-example-elasticsearch-golang/storage"
	"github.com/russmatney/nested-query-example-elasticsearch-golang/trainers"
)

const (
	assetsPath  = "./es-settings.yml"
	indexName   = "poke-index"
	elasticHost = "http://elasticsearch:9200"
)

func main() {
	fmt.Println("Creating store")
	store := storage.NewStore(assetsPath, indexName, elasticHost)

	fmt.Println("Saving mock data to elasticsearch")
	for _, t := range trainers.Trainers {
		if err := store.SaveTrainer(t.Name, t); err != nil {
			panic(err)
		}
	}

}
