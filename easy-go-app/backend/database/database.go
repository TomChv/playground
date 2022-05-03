package database

import (
	_ "github.com/lib/pq"
	"izi-go.com/ent"
)

const dbType = "postgres"

func New(config *config) (*ent.Client, error) {
	return ent.Open(dbType, config.GetURI())
}
