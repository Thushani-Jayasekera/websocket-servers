package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	// Read environment variables
	serviceURL := os.Getenv("SVC_URL")
	choreoApiKey := os.Getenv("CHOREO_API_KEY")

	fmt.Println("Connecting to service:", serviceURL)

	// Configure headers, adding the API key
	headers := http.Header{
		"Choreo-API-Key": []string{choreoApiKey},
	}

	// Establish the WebSocket connection using DefaultDialer
	connectionURL := serviceURL + "/new"
	for {
		// Attempt to connect
		conn, _, err := websocket.DefaultDialer.Dial(connectionURL, headers)
		if err != nil {
			log.Println("Dial error, retrying in 5 seconds:", err)
			time.Sleep(5 * time.Second) // Wait before retrying
			continue
		}
		log.Println("Connected to WebSocket server")

		// Ensure the connection is closed when done
		defer conn.Close()

		// Start the message sending goroutine
		go func(conn *websocket.Conn) {
			ticker := time.NewTicker(2 * time.Second) // Sends a message every 2 seconds
			defer ticker.Stop()
			for range ticker.C {
				message := `{"currency":"EURUSD"}`
				err := conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					log.Println("Write error:", err)
					return
				}
				fmt.Println("Sent:", message)
			}
		}(conn)

		// Listen for incoming messages
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error, reconnecting:", err)
				break // Exit the loop to reconnect
			}
			fmt.Printf("Received: %s\n", message)
		}

		// Wait before retrying the connection
		time.Sleep(5 * time.Second)
		log.Println("Reconnecting to WebSocket server...")
	}

}
