package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/oauth2/clientcredentials"
)

func main() {
	// Read environment variables
	serviceURL := os.Getenv("SVC_URL")
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	tokenUrl := os.Getenv("TOKEN_URL")
	choreoApiKey := os.Getenv("CHOREO_API_KEY")

	// Configure OAuth2 client credentials
	clientCredsConfig := clientcredentials.Config{
		ClientID:     consumerKey,
		ClientSecret: consumerSecret,
		TokenURL:     tokenUrl,
	}
	client := clientCredsConfig.Client(context.Background())

	fmt.Println("Connecting to service:", serviceURL, client)

	// Configure headers, adding the API key
	headers := http.Header{
		"Choreo-API-Key": []string{choreoApiKey},
	}

	// Establish the WebSocket connection using DefaultDialer
	conn, _, err := websocket.DefaultDialer.Dial(serviceURL, headers)
	if err != nil {
		log.Fatal("Dial error:", err)
		return
	}
	defer conn.Close()

	// Function to send messages periodically
	go func() {
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
	}()

	// Listen for incoming messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			return
		}
		fmt.Printf("Received: %s\n", message)
	}
}
