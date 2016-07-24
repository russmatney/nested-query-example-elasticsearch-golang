package trainers

import "fmt"

// Pokemon defines the structure of a trainer's pokemon
type Pokemon struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

// Trainer defines the type and fields for a trainer
type Trainer struct {
	Name    string     `json:"name"`
	Pokemon []*Pokemon `json:"pokemon"`
}

func (t *Trainer) String() string {
	sum := fmt.Sprintf("\n\tTrainer: %s, Pokemon (%d):", t.Name, len(t.Pokemon))
	for _, pk := range t.Pokemon {
		sum = fmt.Sprintf("%s\n\t\t[ Name: %s, Level: %d ]", sum, pk.Name, pk.Level)
	}
	return sum
}

// Trainers is local example data
var Trainers = []*Trainer{
	{
		Name: "Ash",
		Pokemon: []*Pokemon{
			{Name: "Pikachu", Level: 7},
			{Name: "Charizard", Level: 45},
		},
	}, {
		Name: "Brock",
		Pokemon: []*Pokemon{
			{Name: "Goldar", Level: 9000},
		},
	}, {
		Name: "Misty",
		Pokemon: []*Pokemon{
			{Name: "Rubberduck", Level: 1},
			{Name: "Magikarp", Level: 19},
		},
	}, {
		Name: "Fisherman Steve",
		Pokemon: []*Pokemon{
			{Name: "Magikarp", Level: 5},
			{Name: "Magikarp", Level: 5},
			{Name: "Magikarp", Level: 5},
			{Name: "Magikarp", Level: 5},
			{Name: "Psyduck", Level: 19},
		},
	},
}
