import ballerina/io;
import ballerina/websocket;
import ballerina/log;
import ballerina/tcp;

@websocket:ServiceConfig {
   subProtocols: ["chat", "foo"]
}

service /chat on new websocket:Listener(8085) {

     resource function get rocket/[string rocketId]/status() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New sever WebSocket connection");
        return new ChatService();
    }
}

service /chat2 on new websocket:Listener(8086) {
     resource function get rocket/[string rocketId]/status() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New sever WebSocket connection", rocketId);
        return new ChatService();
    }
}

service /chat3 on new websocket:Listener(8087) {
     resource function get rocket/[string rocketId]/status() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New sever WebSocket connection", rocketId);
        return new ChatService();
    }
}

service /chat4 on new websocket:Listener(8088) {
     resource function get rocket/[string rocketId]/status() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New sever WebSocket connection", rocketId);
        return new ChatService();
    }
}

service /chat5 on new websocket:Listener(8089) {
     resource function get rocket/[string rocketId]/status() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New sever WebSocket connection", rocketId);
        return new ChatService();
    }
}

service class ChatService {
    *websocket:Service;

    // This `remote function` is triggered when a new message is received
    // from a client. It accepts `anydata` as the function argument. The received data 
    // will be converted to the data type stated as the function argument.
    remote isolated function onTextMessage(websocket:Caller caller, string text) returns websocket:Error? {
        io:println(text, caller.getConnectionId());
        check caller->writeTextMessage("Hello!, How are you?");
    }

    remote isolated function onBinaryMessage(websocket:Caller caller, byte[] data) returns websocket:Error? {
        io:println("Binary data", data);
        check caller->writeBinaryMessage(data);
    }

    remote function onPing(websocket:Caller caller, byte[] data) returns error? {
        io:println(string `Ping received with data: ${data.toBase64()}`);
        check caller->pong(data);
    }
 
    remote function onPong(websocket:Caller caller, byte[] data) {
        io:println(string `Pong received with data: ${data.toBase64()}`);
    }

    remote function onError(tcp:Error err) returns tcp:Error? {
        log:printError("An error occurred", 'error = err);
    }

    remote function onClose(websocket:Caller caller, int statusCode, string reason) {
        io:println(string `Client closed connection with ${statusCode} because of ${reason} ` , caller.getConnectionId());
    }

}
