const WebSocket = require('ws');
const fs = require('fs');
const path = require('path');

// Path to your video file
const videoFilePath = path.join(__dirname, 'recording.mov');
const CHUNK_SIZE =  50 *1024 * 1024; // 50MB chunks

// Function to stream video over WebSocket
function streamVideo(ws, videoPath) {
  const readStream = fs.createReadStream(videoPath, { highWaterMark: CHUNK_SIZE });

  readStream.on('data', (chunk) => {
    if (ws.readyState === WebSocket.OPEN) {
      console.log(`Sending chunk of size: ${chunk.length}`);
      ws.send(chunk); // Send each chunk to the WebSocket server
    }
  });

  readStream.on('end', () => {
    console.log('Video streaming completed. Restarting video stream...');
    // Restart the video stream after a short delay
    setTimeout(() => streamVideo(ws, videoPath), 0); // 1 second delay
  });

  readStream.on('error', (err) => {
    console.error(`Error reading video file: ${err}`);
    ws.close();
  });
}

// WebSocket server address
const serverUrl = 'wss://1704d9ae-d506-4efb-9982-410b58842006-dev.e1-us-east-azure.st.choreoapis.dev/uuzz/gowsservice/v1.0/new';

// Create WebSocket connection
const ws = new WebSocket(serverUrl);

// On WebSocket connection open
ws.on('open', () => {
  console.log('WebSocket connection established.');
  // Start streaming the video file repeatedly
  streamVideo(ws, videoFilePath);
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
});
