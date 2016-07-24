package trainers

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
	},
}
