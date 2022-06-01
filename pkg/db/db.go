package db

import (
	"database/sql"
	"fmt"

	//"net/url"
	//"flag"

	_ "github.com/lib/pq"
)

type Connection struct {
	DbName         string
	User, Password string
	Host           string
	Port           string
	DisableSSL     bool
}

/* func main() {

	conn := Connection{
		DbName:     "forumsdb",
		User:       "postgres",
		Password:   "qwerty334455",
		Host:       "127.0.0.1",
		Port:       "5050",
		DisableSSL: true,
	}

	conn.Open()
} */

func (c *Connection) Open() (*sql.DB, error) {
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", c.User, c.Password, c.Host, c.Port, c.DbName)
	return sql.Open("postgres", connStr)
}
