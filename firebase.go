package main

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	// アプリの設定ファイル等の指定
	opt := option.WithCredentialsFile("firebase_token.json")
	config := &firebase.Config{
		DatabaseURL: "https://magiot.firebaseio.com/",
	}

	// アプリの初期化
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// クライアント作成
	client, err := app.Database(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	var flags map[string]int

	// リファレンスの取得
	ref := client.NewRef("data/flag")

	// データの取得
	err = ref.Get(context.Background(), &flags)
	if err != nil {
		log.Fatal(err)
	}
	for name, status := range flags {
		fmt.Println(name, status)
	}
}
