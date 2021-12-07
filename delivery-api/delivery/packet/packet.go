package packet

import (
	"encore.app/delivery/database"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db = database.Client()
}
