const WebSocket = require('ws');
const fs = require('fs');
const path = require('path');

// Path to your video file
const videoFilePath = path.join(__dirname, 'recording.dat');
const CHUNK_SIZE = 200 * 1024 * 1024; // 50MB chunks
let ws; // Global WebSocket variable
const RECONNECT_DELAY = 5000; // Delay before reconnecting (in milliseconds)

// Function to stream video over WebSocket
function streamVideo() {
  const readStream = fs.createReadStream(videoFilePath, { highWaterMark: CHUNK_SIZE });

  readStream.on('data', (chunk) => {
    if (ws.readyState === WebSocket.OPEN) {
      console.log(`Sending chunk of size: ${chunk.length}`);
      ws.send(chunk); // Send each chunk to the WebSocket server
    }
  });

  readStream.on('end', () => {
    console.log('Video streaming completed. Restarting video stream...');
    // Restart the video stream after a short delay
    setTimeout(() => streamVideo(), 0); // 1 second delay
  });

  readStream.on('error', (err) => {
    console.error(`Error reading video file: ${err}`);
    ws.close();
  });
}

// Function to connect to WebSocket server
function connectWebSocket() {
  // WebSocket server address
  const serverUrl = 'wss://1704d9ae-d506-4efb-9982-410b58842006-dev.e1-us-east-azure.st.choreoapis.dev/uuzz/gowsservice/v1.0/new';
  ws = new WebSocket(serverUrl);

  // On WebSocket connection open
  ws.on('open', () => {
    console.log('WebSocket connection established.');
    // Start streaming the video file repeatedly
    streamVideo();
  });

  // Handle messages from the server
  ws.on('message', (data) => {
    console.log(`Received message from server: ${data}`);
  });

  // Handle WebSocket errors
  ws.on('error', (error) => {
    console.error(`WebSocket error: ${error}`);
  });

  // Handle WebSocket closure
  ws.on('close', (code, reason) => {
    console.log(`WebSocket connection closed: ${code} - ${reason}`);
    // Attempt to reconnect after a delay
    setTimeout(connectWebSocket, RECONNECT_DELAY);
  });
}

// Start the initial connection
connectWebSocket();
