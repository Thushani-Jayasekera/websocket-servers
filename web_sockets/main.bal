import ballerina/io;
import ballerina/websocket;
import ballerina/log;
import ballerina/tcp;

@websocket:ServiceConfig {
   subProtocols: ["chat", "foo"]
}

service /chat on new websocket:Listener(8085) {

    resource function get rocket/[string rocketId]/status/[string rocketId1]/[string rocketId2]/[string rocketId3]() returns websocket:Service {
        // Accept the WebSocket upgrade by returning a `websocket:Service`.
        io:println("New rocket WebSocket connection", rocketId);
        return new ChatService();
    }
}

service class ChatService {
    *websocket:Service;

    // This `remote function` is triggered when a new message is received
    // from a client. It accepts `anydata` as the function argument. The received data 
    // will be converted to the data type stated as the function argument.
    remote function onMessage(websocket:Caller caller, string chatMessage) returns error? {
        io:println(chatMessage, caller.getConnectionId());
        check caller->writeMessage("Hello!, How are you?");
    }

    remote function onError(tcp:Error err) returns tcp:Error? {
        log:printError("An error occurred", 'error = err);
    }

    remote function onClose(websocket:Caller caller, int statusCode, string reason) {
        io:println(string `Client closed connection with ${statusCode} because of ${reason} ` , caller.getConnectionId());
    }

}
