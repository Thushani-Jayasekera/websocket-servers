import ballerina/io;
import ballerina/websocket;
import ballerina/lang.runtime;
import ballerina/time;

public function main() returns error? {

    time:Utc startTime = time:utcNow();
    // Create a new WebSocket client.
    websocket:Client chatClient = check new ( "ws://websocketproject-3275271155:9090/chat",
        //"wss://3ea87e83-4b86-45d2-8bd3-6a0348ea25d7-dev.e1-us-east-azure.perf.choreoapis.dev/websockettests/websocketperf8/websocket-5cd/v1.0",
        secureSocket = {enable: false} );


    while (true) {
        // Write a message to the server using `writeMessage`.
        // This function accepts `anydata`. If the given type is a `byte[]`, the message will be sent as
        // binary frames and the rest of the data types will be sent as text frames.
        io:println(time:utcDiffSeconds(time:utcNow(), startTime), "Sending message to the server");
        check chatClient->writeMessage("Hello John 10!");
        

        // Read a message sent from the server using `readMessage`.
        // The contextually-expected data type is inferred from the LHS variable type. The received data
        // will be converted to that particular data type.
        io:println("Reading message from the server");
        string message = check chatClient->readMessage();

        io:println(message);
    
        runtime:sleep(5);
    }

}
