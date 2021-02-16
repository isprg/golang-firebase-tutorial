package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	opt := option.WithCredentialsFile("firebase_token.json")
	config := &firebase.Config{
		DatabaseURL: "https://magiot.firebaseio.com/",
	}
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var flags map[string]int
	ref := client.NewRef("data/flag")
	err = ref.Get(context.Background(), &flags)
	if err != nil {
		log.Fatal(err)
	}
	for name, status := range flags {
		fmt.Println(name, status)
	}
}
