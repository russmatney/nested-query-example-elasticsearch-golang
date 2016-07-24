package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/russmatney/nested-query-example-elasticsearch-golang/storage"
	"github.com/russmatney/nested-query-example-elasticsearch-golang/trainers"
)

const (
	assetsPath  = "./es-settings.yml"
	indexName   = "poke-index"
	elasticHost = "http://elasticsearch:9200"
)

func main() {
	log.Info("Creating store")
	store := storage.NewStore(assetsPath, indexName, elasticHost)

	log.Info("Saving mock data to elasticsearch")
	for _, t := range trainers.Trainers {
		if err := store.SaveTrainer(t.Name, t); err != nil {
			panic(err)
		}
	}

	log.Info("Searching for trainers")
	searchForMagikarp(store)
}

func searchForMagikarp(store storage.Storage) {
	magikarpSearch := storage.TrainerSearchOpts{
		Pokemon: "Magikarp",
		Level:   19,
	}
	trainers, _ := store.FetchTrainers(magikarpSearch)

	log.WithFields(log.Fields{
		"search":   magikarpSearch,
		"trainers": trainers,
	}).Info("Trainers found for search")
}
