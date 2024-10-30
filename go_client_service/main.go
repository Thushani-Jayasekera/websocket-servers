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
	// consumerKey := os.Getenv("CONSUMER_KEY")
	// consumerSecret := os.Getenv("CONSUMER_SECRET")
	// tokenUrl := os.Getenv("TOKEN_URL")
	choreoApiKey := os.Getenv("CHOREO_API_KEY")

	// Configure OAuth2 client credentials
	// clientCredsConfig := clientcredentials.Config{
	// 	ClientID:     consumerKey,
	// 	ClientSecret: consumerSecret,
	// 	TokenURL:     tokenUrl,
	// }
	// client := clientCredsConfig.Client(context.Background())

	fmt.Println("Connecting to service:", serviceURL, choreoApiKey)

	// Configure headers, adding the API key
	headers := http.Header{
		"Choreo-API-Key": []string{choreoApiKey},
		// "Test-Key":       []string{"eyJraWQiOiJnYXRld2F5X2NlcnRpZmljYXRlX2FsaWFzIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiJhMzllYTQxYS00MzgyLTQ3ZWQtODRiZi01NjM3NjE2ODk1NWRAY2FyYm9uLnN1cGVyIiwiYXVkIjoiY2hvcmVvOmRlcGxveW1lbnQ6c2FuZGJveCIsImlzcyI6Imh0dHBzOlwvXC9zdHMucHJldmlldy1kdi5jaG9yZW8uZGV2OjQ0M1wvYXBpXC9hbVwvcHVibGlzaGVyXC92MlwvYXBpc1wvaW50ZXJuYWwta2V5Iiwia2V5dHlwZSI6IlNBTkRCT1giLCJzdWJzY3JpYmVkQVBJcyI6W3sic3Vic2NyaWJlclRlbmFudERvbWFpbiI6bnVsbCwibmFtZSI6ImNvbXBvbmVudHlhbWwgLSBnb3dlYnNvY2tldGFwaSIsImNvbnRleHQiOiJcL2YzNGZjMDExLTM4YTYtNDY2My05ZmExLTlhMjQzMDA2YzUzOFwvd2Vic29ja2V0cHJvamVjdFwvY29tcG9uZW50eWFtbFwvdjEuMCIsInB1Ymxpc2hlciI6ImNob3Jlb19kZXZfYXBpbV9hZG1pbiIsInZlcnNpb24iOiJ2MS4wIiwic3Vic2NyaXB0aW9uVGllciI6bnVsbH1dLCJleHAiOjE3MzAyODg0MzQsInRva2VuX3R5cGUiOiJJbnRlcm5hbEtleSIsImlhdCI6MTczMDI4NzgzNCwianRpIjoiZWY3MDAzNzQtNGNlYi00Y2ExLWEyMjQtNWY2ZDRjODI2NjM3In0.hKjECY8zmxGkl5vVidg1uCfNgT08DJxrRcQl78zQ8tidgGXdrmPTNNhXkvP8t2DI_2op0jn7mEL_yZd9BzjkbOutHlkvNovhLZtAsmAc0LkJt529XTsxJhbnlz2wcrMFw1gMS2iRt2nSjUYAYVTZYBU4zq537G1vlOXCydagNjCzFdYuTMoQJbjdA--WSO6WzpgIg2GkJ1k14UXdjPxuxNTr014m5o5agkiQPGvDHoZgjk_r3lgCxr9ruDd0G5mm4hG0f7HoxAsrRTQNHuSttsrtsb-RpF_4B7ucmuAv5XZePsigkJXfxbHAJybK5pZQx1DZ4LvYLOwo8Klan_QAYMyIv-5WCTvc-vMsprrXQCFgr1MVygbxTKOk0bbt4tVm8JPhwyWwvHlU6xUqMeap5aWLq5OPYVg0E_vNRyezftyY_TA2HxuZ4CrUjaXG4W-bUJImf9NvNyDvrzA0v1Ph2ZUeYdoKi9VNPQMNQt2wbIQKZgVVdkaSKxUBuiDpuaXkc7rDBziZSZ97iJ9Nj8jrRpRrObG3MfCcAFhG4uV36mWXDGm5mBZhtKiEPPscye3GtCoCKF6RYf5BuR3PnrVe9EY_r7YOve8lcj0RWtiTmMnfbR0PZJKbz6rh2WJYSgox2BnXW2D2YARk6xilyt62yvoTGjFrpVD1kQFuQPYQdyo"},
	}

	// Establish the WebSocket connection using DefaultDialer
	connectionURL := serviceURL + "/echo"
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
