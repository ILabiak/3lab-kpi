package main

import (
	//"database/sql"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/ILabiak/3lab-kpi/pkg/db"
	"github.com/ILabiak/3lab-kpi/pkg/forums"
)

var httpPortNumber = flag.Int("p", 8080, "HTTP port number")

func NewDbConnection() (*sql.DB, error) {
	conn := db.Connection{
		DbName:     "forumsdb",
		User:       "postgres",
		Password:   "qwerty334455",
		Host:       "127.0.0.1",
		Port:       "5050",
		DisableSSL: true,
	}
	return conn.Open()
}

func main() { //testing database functions
	flag.Parse()

	// db, err := NewDbConnection()
	// if err != nil {
	// 	fmt.Printf("%v", err)
	// }
	// data := forums.NewData(db)
	// interests := []uint8("politics,literature")
	// err = forums.CreateUser(data, "testuser", interests)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	if server, err := ComposeApiServer(HttpPortNumber(*httpPortNumber)); err == nil {
		go func() {
			log.Println("Starting forum server...")

			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("HTTP server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}
		}()
		sigChannel := make(chan os.Signal, 1)
		signal.Notify(sigChannel, os.Interrupt)
		<-sigChannel

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		}
	} else {
		log.Fatalf("Cannot initialize forum server: %s", err)
	}

}

func getArrString(arr []*forums.Forum) string {
	res := ``
	for _, obj := range arr {
		str := fmt.Sprintf(`{Id:%d,
Name:%s,
TopicKeyword:%s,
Users:%s}

`, obj.Id, obj.Name, obj.TopicKeyword, string(obj.Users))
		res += str
	}
	return res
}
