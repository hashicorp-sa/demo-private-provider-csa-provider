package provider

import (
	"strings"
	"time"
)

type csa_client struct {
	id string
}

func (c *csa_client) GetSpecialism() string {
	specialisms := make(map[string]string)
	specialisms[""] = ""
	specialisms["jared holgate"] = "Terraform"
	specialisms["jan repnak"] = "Consul"
	specialisms["colin turney"] = "Vault"

	return specialisms[strings.ToLower(c.id)]
}

func (c *csa_client) GetSetupDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
