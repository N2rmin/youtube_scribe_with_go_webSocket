package main

import (
	"fmt"
	"log"
	"net/http"
	"youtubemonitorwebsocket/websocket"
)

func homepage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./index.html")

}

func stats(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)

	}
	go websocket.Writer(ws)
}

func setupRoutes() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/stats", stats)

	log.Fatal(http.ListenAndServe(":8085", nil))
}

func main() {
	fmt.Println("Youtube Subcriber Monitor")
	setupRoutes()

}
