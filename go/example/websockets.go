package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// CheckOrigin can be added if you need to handle CORS.
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust this according to your CORS policy
	},
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade HTTP server connection to the WebSocket protocol
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("error upgrading connection: %v", err)
			return
		}
		defer conn.Close() // Ensure the connection is closed when the function returns

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				log.Printf("error reading message: %v", err)
				break // Exit the loop if there's an error
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			// Write message back to browser
			if err = conn.WriteMessage(msgType, msg); err != nil {
				log.Printf("error writing message: %v", err)
				break // Exit the loop if there's an error
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	// Start the HTTP server on port 8009
	log.Println("Server started on port 8009")
	err := http.ListenAndServe(":8009", nil)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}
}
