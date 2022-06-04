package forums

import (
	//"encoding/json"
	"encoding/json"
	"log"
	"net/http"

	"github.com/ILabiak/3lab-kpi/pkg/tools"
)

type HttpHandlerFunc http.HandlerFunc

func HttpHandler(data *Data) HttpHandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			handleListForums(data, rw)
			//} else if r.Method == "POST" {
			//handleUserCreate(r, rw, data)
		} else {
			rw.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func handleUserCreate(r *http.Request, rw http.ResponseWriter, data *Data) {
	var f Forum
	if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
		log.Printf("Error decoding forum input: %s", err)
		tools.WriteJsonBadRequest(rw, "bad JSON payload")
		return
	}
	err := CreateForum(data, f.Name, f.TopicKeyword, f.Users)
	if err == nil {
		tools.WriteJsonOk(rw, &f)
	} else {
		log.Printf("Error inserting record: %s", err)
		tools.WriteJsonInternalError(rw)
	}
}

func handleListForums(data *Data, rw http.ResponseWriter) {
	res, err := ListForums(data)
	if err != nil {
		log.Printf("Error making query to the db: %s", err)
		tools.WriteJsonInternalError(rw)
		return
	}
	tools.WriteJsonOk(rw, res)
}
