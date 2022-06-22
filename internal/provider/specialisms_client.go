package provider

import (
	"strings"
	"time"
)

type animal_client struct {
	id string
}

func (c *animal_client) GetAnimalFromClass() string {
	animals := make(map[string]string)
	animals[""] = "Duck Billed Platipus"
	animals["mammal"] = "Horse"
	animals["bird"] = "Peregrine Falcon"
	animals["invertebrate"] = "Stag Beetle"
	animals["fish"] = "Great White Shark"
	animals["reptile"] = "Blue Iguana"
	animals["amphibian"] = "Common Frog"

	return animals[strings.ToLower(c.id)]
}

func (c *animal_client) GetSetupDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
