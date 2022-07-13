package animals

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetAnimalFromClass(t *testing.T) {
	var animal Animal

	classes := []string{"MaMMal", "biRD", "", "inverteBRAte", "fiSH", "Reptile", "amphibian"}
	animalExamples := []string{"Horse", "Peregrine Falcon", "Duck Billed Platipus", "Stag Beetle", "Great White Shark", "Blue Iguana", "Common Frog"}

	for i, class := range classes {
		example := animal.GetAnimalFromClass(class)
		assert.Equal(t, animalExamples[i], example)
	}
}

func TestGetAnimalFromClassNoExists(t *testing.T) {
	var animal Animal

	example := animal.GetAnimalFromClass("dfsgdsfgdfsgsg")
	assert.Equal(t, "", example)
}

func TestGetSetupDate(t *testing.T) {
	var animal Animal

	example := animal.GetSetupDate()
	_, err := time.Parse("2006-01-02 15:04:05", example)
	assert.NoError(t, err)
}
