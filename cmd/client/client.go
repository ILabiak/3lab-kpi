package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/ILabiak/3lab-kpi/pkg/sdk"
)

var (
	target  = flag.String("target", "http://localhost:8080", "Target server address")
	timeout = flag.Duration("timeout", 5*time.Second, "Timeout for scenarios execution")
)

func main() {
	flag.Parse()
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	client := &sdk.Client{BaseUrl: *target}

	fmt.Println("=== Scenario 1 ===")
	forums, err := client.ListForums(ctx)
	if err != nil {
		log.Fatal("Cannot list forums: ", err)
	}
	fmt.Print("Available forums: ")
	fmt.Println(forums)

	fmt.Println("=== Scenario 2 ===")
	err = client.CreateUser(ctx, "new_user")
	if err != nil {
		log.Fatal("Cannot create user: ", err)
	}
	fmt.Println("Created a new user")
	forums, err = client.ListForums(ctx)
	if err != nil {
		log.Fatal("Cannot list forums: ", err)
	}
	fmt.Print("New forums list: ")
	fmt.Println(forums)
}
