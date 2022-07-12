package animals

import (
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

func (c *Client) GetAnimalFromClass(id string) string {
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

func (c *Client) GetSetupDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
