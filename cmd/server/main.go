package main

import (
	//"database/sql"
	"fmt"

	"github.com/ILabiak/3lab-kpi/pkg/db"
	"github.com/ILabiak/3lab-kpi/pkg/forums"
)

func main() { //testing database functions
	conn := db.Connection{
		DbName:     "forumsdb",
		User:       "postgres",
		Password:   "qwerty334455",
		Host:       "127.0.0.1",
		Port:       "5050",
		DisableSSL: true,
	}
	db, err := conn.Open()
	if err != nil {
		fmt.Printf("%v", err)
	}
	data := forums.NewData(db)
	/* 	res, err := forums.ListForums(data)
	   	if err != nil {
	   		fmt.Printf("%v", err)
	   	}
	   	str := getArrString(res) */
	//users :=[]uint8("user1,user3")
	//str := []string{"user1", "user3"}
	interests := []uint8("politics,literature")
	//err = forums.CreateForum(data, "lol", "kek", users)
	err = forums.CreateUser(data, "testuser", interests)
	if err != nil {
		fmt.Println(err)
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
