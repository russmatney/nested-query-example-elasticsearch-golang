package storage

import (
	"io/ioutil"

	log "github.com/Sirupsen/logrus"

	"github.com/ghodss/yaml"
	"github.com/olivere/elastic"
	trainers "github.com/russmatney/nested-query-example-elasticsearch-golang/trainers"
)

// Storage is the exposed interface for use in other packages
type Storage interface {
	FetchTrainers(TrainerSearchOpts) ([]*trainers.Trainer, error)
	SaveTrainer(string, *trainers.Trainer) error
}

var _ Storage = &Store{}

// Store implements Storage
type Store struct {
	es           *elastic.Client
	trainerIndex string
	trainerType  string
}

// NewStore creates and returns a new store
func NewStore(assetsPath, indexName, esURL string) Storage {
	client, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		log.WithError(err).Fatal("Failed to start elasticsearch client")
	}
	store := &Store{
		es:           client,
		trainerIndex: indexName,
		trainerType:  "trainer",
	}

	if err := store.bootstrap(assetsPath, indexName); err != nil {
		log.WithError(err).Fatal("Failed to bootstrap elasticsearch store")
	}
	return store
}

// Bootstrap applies settings to the elasticsearch index
func (s *Store) bootstrap(settingsPath, indexName string) error {
	exists, err := s.es.IndexExists(indexName).Do()
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	settings, err := ioutil.ReadFile(settingsPath)
	if err != nil {
		return err
	}
	json, err := yaml.YAMLToJSON([]byte(settings))
	if err != nil {
		return err
	}
	_, err = s.es.CreateIndex(indexName).BodyString(string(json)).Do()
	if err != nil {
		return err
	}

	return nil
}

// SaveTrainer indexes a passed trainer object
func (s *Store) SaveTrainer(id string, trainer *trainers.Trainer) error {
	_, err := s.es.Index().Index(s.trainerIndex).Type(s.trainerType).Id(id).BodyJson(trainer).Do()
	if err != nil {
		return err
	}
	return nil
}
