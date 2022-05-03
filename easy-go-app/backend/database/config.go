package database

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	Username string `envconfig:"DB_USER" required:"true"`
	Name     string `envconfig:"DB_NAME" required:"true"`
	Password string `envconfig:"DB_PASS" required:"true"`
}

var Config config

func (c *config) GetURI() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		c.Host,
		c.Port,
		c.Username,
		c.Name,
		c.Password)
}

func init() {
	if err := envconfig.Process("", &Config); err != nil {
		log.Fatal(err)
	}
}
