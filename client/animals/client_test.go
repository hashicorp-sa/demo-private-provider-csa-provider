package animals

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const url = "https://test.com"
const token = "4353424g345425b23b434b345b"

func TestEmptyClient(t *testing.T) {
	client, _ := New("", "")

	assert.Equal(t, "", client.url)
	assert.Equal(t, "", client.token)
}

func TestNewClient(t *testing.T) {
	client, _ := New(url, token)

	assert.Equal(t, url, client.url)
	assert.Equal(t, token, client.token)
}

func TestCreateAnimal(t *testing.T) {
	client, _ := New(url, token)
	animalCreate := AnimalCreateModel{ Class: "Bird", }

	animal, _ := client.Create(animalCreate)

	assert.Equal(t, "Bird", animal.Class)
	assert.Equal(t, "Peregrine Falcon", animal.Animal)
	_, err := uuid.Parse(animal.Id)
	assert.NoError(t, err)
	_, err = time.Parse("2006-01-02 15:04:05", animal.Created)
	assert.NoError(t, err)
}

func TestUpdateAnimal(t *testing.T) {
	client, _ := New(url, token)
	animalUpdate := AnimalUpdateModel{ 
		Id: uuid.New().String(),  
		Class: "Bird",  
	}

	animal, _ := client.Update(animalUpdate)

	assert.Equal(t, "Bird", animal.Class)
	assert.Equal(t, "Peregrine Falcon", animal.Animal)
	_, err := uuid.Parse(animal.Id)
	assert.NoError(t, err)
	_, err = time.Parse("2006-01-02 15:04:05", animal.Created)
	assert.NoError(t, err)
}

func TestReadAnimal(t *testing.T) {
	client, _ := New(url, token)
	animalRead := AnimalReadModel{ 
		Id: uuid.New().String(),  
		Class: "Bird", 
		Created: "2006-01-02 15:04:05", 
	}

	animal, _ := client.Read(animalRead)

	assert.Equal(t, "Bird", animal.Class)
	assert.Equal(t, "Peregrine Falcon", animal.Animal)
	_, err := uuid.Parse(animal.Id)
	assert.NoError(t, err)
	_, err = time.Parse("2006-01-02 15:04:05", animal.Created)
	assert.NoError(t, err)
}

func TestDeleteAnimal(t *testing.T) {
	client, _ := New(url, token)
	animalDelete := AnimalDeleteModel{ 
		Id: uuid.New().String(),   
	}

	err := client.Delete(animalDelete)

	assert.NoError(t, err)
}