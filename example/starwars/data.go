package starwars

import graphql "github.com/neelance/graphql-go"

var humans = []*human{
	{
		ID:        "1000",
		Name:      "Luke Skywalker",
		Friends:   []graphql.ID{"1002", "1003", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.72,
		Mass:      77,
		Starships: []graphql.ID{"3001", "3003"},
	},
	{
		ID:        "1001",
		Name:      "Darth Vader",
		Friends:   []graphql.ID{"1004"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    2.02,
		Mass:      136,
		Starships: []graphql.ID{"3002"},
	},
	{
		ID:        "1002",
		Name:      "Han Solo",
		Friends:   []graphql.ID{"1000", "1003", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.8,
		Mass:      80,
		Starships: []graphql.ID{"3000", "3003"},
	},
	{
		ID:        "1003",
		Name:      "Leia Organa",
		Friends:   []graphql.ID{"1000", "1002", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.5,
		Mass:      49,
	},
	{
		ID:        "1004",
		Name:      "Wilhuff Tarkin",
		Friends:   []graphql.ID{"1001"},
		AppearsIn: []string{"NEWHOPE"},
		Height:    1.8,
		Mass:      0,
	},
}

var humanData = make(map[graphql.ID]*human)

func init() {
	for _, h := range humans {
		humanData[h.ID] = h
	}
}

var droids = []*droid{
	{
		ID:              "2000",
		Name:            "C-3PO",
		Friends:         []graphql.ID{"1000", "1002", "1003", "2001"},
		AppearsIn:       []string{"NEWHOPE", "EMPIRE", "JEDI"},
		PrimaryFunction: "Protocol",
	},
	{
		ID:              "2001",
		Name:            "R2-D2",
		Friends:         []graphql.ID{"1000", "1002", "1003"},
		AppearsIn:       []string{"NEWHOPE", "EMPIRE", "JEDI"},
		PrimaryFunction: "Astromech",
	},
}

var droidData = make(map[graphql.ID]*droid)

func init() {
	for _, d := range droids {
		droidData[d.ID] = d
	}
}

var starships = []*starship{
	{
		ID:     "3000",
		Name:   "Millennium Falcon",
		Length: 34.37,
	},
	{
		ID:     "3001",
		Name:   "X-Wing",
		Length: 12.5,
	},
	{
		ID:     "3002",
		Name:   "TIE Advanced x1",
		Length: 9.2,
	},
	{
		ID:     "3003",
		Name:   "Imperial shuttle",
		Length: 20,
	},
}

var starshipData = make(map[graphql.ID]*starship)

func init() {
	for _, s := range starships {
		starshipData[s.ID] = s
	}
}

var reviews = make(map[string][]*review)
