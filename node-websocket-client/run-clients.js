const { spawn } = require('child_process');

const maxClients = 10; // Number of clients
const retries = {}; // To keep track of retries for each client
const maxRetries = 10; // Maximum number of retries per client

// Function to run a client
function runClient(clientId) {
    const clientProcess = spawn('node', ['index.js']);

    // Listen for stdout and log the output
    clientProcess.stdout.on('data', (data) => {
        console.log(`Client ${clientId}: ${data}`);
    });

    // Listen for stderr and log errors
    clientProcess.stderr.on('data', (data) => {
        console.error(`Client ${clientId} error: ${data}`);
    });

    // Handle process exit
    clientProcess.on('exit', (code, signal) => {
        console.log(`Client ${clientId} exited with code ${code} and signal ${signal}`);
        if (code === 503) {  // Assuming your index.js exits with code 503 on 503 error
            retries[clientId] = (retries[clientId] || 0) + 1;
            if (retries[clientId] < maxRetries) {
                console.log(`Retrying client ${clientId}...`);
                runClient(clientId); // Retry the client
            } else {
                console.error(`Client ${clientId} exceeded max retries`);
            }
        }
    });
}

// Create 10 clients
for (let i = 0; i < maxClients; i++) {
    runClient(i + 1);
}
