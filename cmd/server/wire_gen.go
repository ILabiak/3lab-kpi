package main

import "github.com/ILabiak/3lab-kpi/pkg/forums"

func ComposeApiServer(port HttpPortNumber) (*ForumsApiServer, error) {
	db, err := NewDbConnection()
	if err != nil {
		return nil, err
	}
	store := forums.NewData(db)
	httpHandlerFunc := forums.HttpHandler(store)
	forumsApiServer := &ForumsApiServer{
		Port:          port,
		ForumsHandler: httpHandlerFunc,
	}
	return forumsApiServer, nil
}
