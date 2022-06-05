package forums

import (
	//"encoding/json"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ILabiak/3lab-kpi/pkg/tools"
)

type HttpHandlerFunc http.HandlerFunc

type UserDto struct {
	username  string
	interests []uint8
}

type ForumOutput struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	TopicKeyword string  `json:"topickeyword"`
	Users        string `json:"users"`
}

func HttpHandler(data *Data) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(data, rw)
		} else if r.Method == "POST" {
			handleUserCreate(r, rw, data)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, data *Data) {
	var u UserDto
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		log.Printf("Error decoding forum input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := CreateUser(data, u.username, u.interests)
	if err == nil {
		tools.WriteJsonOk(rw, &u)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListForums(data *Data, rw http.ResponseWriter) {
	var modifiedRes []*ForumOutput
	res, err := ListForums(data)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	for _, el := range res {
		obj := ForumOutput{
			Id : el.Id,
			Name :el.Name,
			TopicKeyword: el.TopicKeyword,
			Users: string(el.Users),
		}
		modifiedRes = append(modifiedRes, &obj)
	}
	tools.WriteJsonOk(rw, modifiedRes)
}
