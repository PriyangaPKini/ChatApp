package main

import (
    "fmt"
    "net/http"

    "github.com/PriyangaPKini/ChatApp/pkg/websocket"
)


func serveWs(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w,r)
    if err != nil {
        fmt.Fprintf(w, "%+V\n", err)
    }
    go websocket.Writer(ws)
    websocket.Reader(ws)
}

func setupRoutes() {
  // map our `/ws` endpoint to the `serveWs` function
    http.HandleFunc("/ws", serveWs)
}

func main() {
    fmt.Println("Distributed Chat App v0.01")
    setupRoutes()
    http.ListenAndServe(":8080", nil)
}
