package animals

import (
	"github.com/google/uuid"
	"strings"
	"time"
)

type Client struct {
	url   string
	token string
}

func New(url string, token string) (client Client, err error) {
	client.url = url
	client.token = token
	return client, nil
}

type AnimalCreateModel struct {
	Class string
}

type AnimalUpdateModel struct {
	Id    string
	Class string
}

type AnimalReadModel struct {
	Id      string
	Class   string
	Created string
}

type AnimalDeleteModel struct {
	Id string
}

func (c *Client) Create(animalCreate AnimalCreateModel) (animal Animal, err error) {
	animal.Id = uuid.New().String()
	animal.Class = animalCreate.Class
	animal.Animal = getAnimalFromClass(animalCreate.Class)
	animal.Created = getSetupDate()
	return animal, nil
}

func (c *Client) Update(animalUpdate AnimalUpdateModel) (animal Animal, err error) {
	animal.Id = animalUpdate.Id
	animal.Class = animalUpdate.Class
	animal.Animal = getAnimalFromClass(animalUpdate.Class)
	animal.Created = getSetupDate()
	return animal, nil
}

func (c *Client) Read(animalRead AnimalReadModel) (animal Animal, err error) {
	animal.Id = animalRead.Id
	animal.Class = animalRead.Class
	animal.Animal = getAnimalFromClass(animalRead.Class)
	animal.Created = animalRead.Created
	return animal, nil
}

func (c *Client) Delete(animalDelete AnimalDeleteModel) (err error) {
	return nil
}

func getAnimalFromClass(id string) string {
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

func getSetupDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
