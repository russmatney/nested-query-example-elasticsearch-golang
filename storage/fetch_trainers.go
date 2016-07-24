package storage

import (
	"github.com/olivere/elastic"
	trainers "github.com/russmatney/nested-query-example-elasticsearch-golang/trainers"
	"reflect"
)

//////////////////////////////////////////////////////////
// Trainer search
//////////////////////////////////////////////////////////

// TrainerSearchOpts is an exposed object to help build a search query
type TrainerSearchOpts struct {
	Pokemon string
	Level   int
}

// FetchTrainers queries trainers with the passed options
func (s *Store) FetchTrainers(ops TrainerSearchOpts) ([]*trainers.Trainer, error) {
	query := elastic.NewBoolQuery()
	query = query.Filter(buildFilters(ops)...)

	res, err := s.es.Search(s.trainerIndex).Type(s.trainerType).Query(query).Do()
	if err != nil {
		return nil, err
	}

	var trns []*trainers.Trainer
	for _, iT := range res.Each(reflect.TypeOf(&trainers.Trainer{})) {
		trns = append(trns, iT.(*trainers.Trainer))
	}
	return trns, nil
}

func buildFilters(opts TrainerSearchOpts) []elastic.Query {
	return []elastic.Query{
		elastic.NewNestedQuery("pokemon", buildPokemonFilter(opts)),
	}
}

func buildPokemonFilter(opts TrainerSearchOpts) elastic.Query {
	return elastic.NewBoolQuery().Filter(
		elastic.NewMatchQuery("pokemon.name", opts.Pokemon),
		elastic.NewMatchQuery("pokemon.level", opts.Level),
	)
}
