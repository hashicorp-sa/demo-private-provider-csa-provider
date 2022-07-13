package animals

import (
	"strings"
	"time"
)

type Animal struct {
	Id      string
	Class   string
	Animal  string
	Created string
}

func (animal *Animal) GetAnimalFromClass(id string) string {
	animals := make(map[string]string)
	animals[""] = "Duck Billed Platipus"
	animals["mammal"] = "Horse"
	animals["bird"] = "Peregrine Falcon"
	animals["invertebrate"] = "Stag Beetle"
	animals["fish"] = "Great White Shark"
	animals["reptile"] = "Blue Iguana"
	animals["amphibian"] = "Common Frog"

	return animals[strings.ToLower(id)]
}

func (animal *Animal) GetSetupDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
