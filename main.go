package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jhersa/chat/client"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
		uri  = host + ":" + port
	)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWs(w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Print("Iniciando server en: ", uri)

	if err := http.ListenAndServe(uri, nil); err != nil {
		log.Fatal(err)
	}
}
