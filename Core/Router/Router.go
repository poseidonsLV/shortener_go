package Router

import (
	"database/sql"
	"playground/MainUrlShortener/Core/Routes/Shortener"
)

func Routes(db *sql.DB) {
	Shortener.Route(db)
}
