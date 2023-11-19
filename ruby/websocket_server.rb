require 'em-websocket'
require 'logger'

PORT = 8009
logger = Logger.new(STDOUT)

EventMachine.run do
  EventMachine::WebSocket.start(host: 'localhost', port: PORT) do |ws|
    ws.onopen do
      logger.info("Client connected")
    end

    ws.onmessage do |msg|
      logger.info("Received message: #{msg}")
      ws.send "Pong: #{msg}"
    end

    ws.onclose do
      logger.info("Client disconnected")
    end

    ws.onerror do |error|
      logger.error("Error occurred: #{error.message}")
    end
  end

  logger.info("WebSocket server started on port #{PORT}")
end
