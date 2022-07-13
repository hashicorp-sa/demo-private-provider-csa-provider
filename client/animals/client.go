package animals

import (
	"github.com/google/uuid"
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
	animal.Animal = animal.GetAnimalFromClass(animalCreate.Class)
	animal.Created = animal.GetSetupDate()
	return animal, nil
}

func (c *Client) Update(animalUpdate AnimalUpdateModel) (animal Animal, err error) {
	animal.Id = animalUpdate.Id
	animal.Class = animalUpdate.Class
	animal.Animal = animal.GetAnimalFromClass(animalUpdate.Class)
	animal.Created = animal.GetSetupDate()
	return animal, nil
}

func (c *Client) Read(animalRead AnimalReadModel) (animal Animal, err error) {
	animal.Id = animalRead.Id
	animal.Class = animalRead.Class
	animal.Animal = animal.GetAnimalFromClass(animalRead.Class)
	animal.Created = animalRead.Created
	return animal, nil
}

func (c *Client) Delete(animalDelete AnimalDeleteModel) (err error) {
	return nil
}
