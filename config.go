package arukas

import (
	API "github.com/arukasio/cli"
)

type Config struct {
}

func (c *Config) NewClient() (*API.Client, error) {
	client, err := API.NewClient()
	if err != nil {
		return nil, err
	}
	client.UserAgent = "Terraform for Arukas"
	return client, nil
}
