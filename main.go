package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WsServer struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewwebsocketServer() *WsServer {
	return &WsServer{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (server *WsServer) Run() {
	for {
		select {
		case client := <-server.register:
			server.registerClient(client)
		case client := <-server.unregister:
			server.unregisterClient(client)
		case message := <-server.broadcast:
			server.broadcastToClients(message)
		}
	}
}

func (server *WsServer) registerClient(client *Client) {
	server.clients[client] = true
}

func (server *WsServer) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}

func (server *WsServer) broadcastToClients(message []byte) {
	for client := range server.clients {
		client.send <- message
	}
}

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

	wsServer := NewwebsocketServer()
	go wsServer.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ServeWs(wsServer, w, r)
	})

	http.Handle("/", http.FileServer(http.Dir("./public")))

	log.Print("Iniciando server en: ", uri)

	if err := http.ListenAndServe(uri, nil); err != nil {
		log.Fatal(err)
	}
}
