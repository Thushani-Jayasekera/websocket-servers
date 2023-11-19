<?php
require __DIR__ . '/vendor/autoload.php';
use Ratchet\Server\IoServer;
use Ratchet\Http\HttpServer;
use Ratchet\WebSocket\WsServer;
use Ratchet\ConnectionInterface;
use Ratchet\MessageComponentInterface;

class EchoServer implements MessageComponentInterface {
    public function onOpen(ConnectionInterface $conn) {
        echo "A client just connected\n";
    }

    public function onMessage(ConnectionInterface $from, $msg) {
        echo sprintf("Received message from client: %s\n", $msg);
        $from->send(sprintf("Pong: %s", $msg));
    }

    public function onClose(ConnectionInterface $conn) {
        echo "A client just disconnected\n";
    }

    public function onError(ConnectionInterface $conn, \Exception $e) {
        echo sprintf("An error has occurred: %s\n", $e->getMessage());
        $conn->close();
    }
}

$server = IoServer::factory(
    new HttpServer(
        new WsServer(
            new EchoServer()
        )
    ),
    8009
);

$server->run();
