const WebSocket = require('ws');
const http = require('http');
const fs = require('fs');
const url = require('url');

// Create a simple HTTP server
const server = http.createServer((req, res) => {
    console.log("creating http server", JSON.stringify(request.headers));
    if (req.url === '/') {
        fs.readFile('websockets.html', (err, data) => {
            if (err) {
                res.writeHead(500);
                return res.end('Error loading websockets.html');
            }
            res.writeHead(200);
            res.end(data);
        });
    }
});

// Create a WebSocket server detached from the HTTP server.
const wss = new WebSocket.Server({ noServer: true });

wss.on('connection', function connection(ws) {
    console.log("A client connected");

    ws.on('message', function incoming(message) {
        console.log('received: %s', message);

        // Echo the message back
        ws.send(message);
    });

    ws.on('close', function close() {
        console.log('Client disconnected');
    });
});

// Handle upgrade manually for '/echo' path
server.on('upgrade', function upgrade(request, socket, head) {
    console.log("on upgrade", JSON.stringify(request.headers));
    const pathname = url.parse(request.url).pathname;

    if (pathname === '/echo') {
        wss.handleUpgrade(request, socket, head, function done(ws) {
            wss.emit('connection', ws, request);
        });
    } else {
        socket.destroy();
    }
});

// Start the server
server.listen(9090, () => {
    console.log('Server is listening on http://localhost:9090');
});
