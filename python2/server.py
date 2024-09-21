import asyncio
import websockets
import logging

logging.basicConfig(level=logging.INFO)

async def echo(websocket, path):
    async for message in websocket:
        logging.info(f"Received message: {message}")
        await websocket.send(message)

start_server = websockets.serve(echo, "0.0.0.0", 8009)

async def main():
    server = await start_server
    logging.info("Server started on port 8009")
    await server.wait_closed()

asyncio.run(main())
