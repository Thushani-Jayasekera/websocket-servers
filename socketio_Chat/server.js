const express = require('express');
const app = express();
const http = require('http');
const server = http.createServer(app);
const { Server } = require("socket.io");

const io = new Server(server,{
              transports: ['websocket'],
              // path: 'ws/api',
              allowUpgrades: true,
              upgradeTimeout: 30000,
});

app.get('/', (req, res) => {
  console.log("sending index.html");
});

io.on('connection', (socket) => {
    console.log('a user connected');

    socket.on('disconnect', () => {
      console.log('user disconnected');
    });

    socket.on('chat message', (msg) => {
        io.emit('chat message', msg);
        console.log('message: ' + msg);
      });

    socket.onAny((event, ...args) => {
        console.log(event, args);
        io.emit('chat message', "received something");
      });
  });

  io.engine.on("connection_error", (err) => {
    console.log(err.req);      // the request object
    console.log(err.code);     // the error code, for example 1
    console.log(err.message);  // the error message, for example "Session ID unknown"
    console.log(err.context);  // some additional error context
  });

server.listen(9090, () => {
  console.log('listening on *:9090');
});
