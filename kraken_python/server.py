import asyncio
import json
import websockets
from datetime import datetime

# Sample data to publish (snapshot and updates)
order_book_snapshot = {
    "channel": "book",
    "type": "snapshot",
    "data": {
        "asks": [
            {"price": 34726.4, "qty": 0.25},
            {"price": 34727.7, "qty": 0.21},
            {"price": 34727.8, "qty": 0.05},
        ],
        "bids": [
            {"price": 34717.6, "qty": 0.13},
            {"price": 34714.1, "qty": 0.08},
        ],
        "checksum": 2645840827,
        "symbol": "BTC/USD"
    }
}

order_book_update = {
    "channel": "book",
    "type": "update",
    "data": {
        "asks": [
            {"price": 34726.4, "qty": 0},
            {"price": 34739.7, "qty": 2.1541},
        ],
        "bids": [],
        "checksum": 4022926185,
        "symbol": "BTC/USD",
        "timestamp": datetime.utcnow().isoformat() + 'Z'
    }
}

# Connected clients
connected_clients = set()

async def subscribe(ws, path):
    """Handles client subscriptions and data publishing"""
    async for message in ws:
        try:
            data = json.loads(message)
            # Subscription request from client
            if data.get("method") == "subscribe":
                symbol = data["params"]["symbol"]
                req_id = data.get("req_id", None)
                
                # Respond with successful subscription
                subscribe_response = {
                    "method": "subscribe",
                    "result": {
                        "channel": "book",
                        "symbol": symbol[0],
                        "depth": 10,
                        "snapshot": True
                    },
                    "success": True,
                    "time_in": datetime.utcnow().isoformat() + 'Z',
                    "time_out": datetime.utcnow().isoformat() + 'Z',
                    "req_id": req_id
                }
                await ws.send(json.dumps(subscribe_response))

                # Send initial snapshot data
                await ws.send(json.dumps(order_book_snapshot))

                # Add to connected clients
                connected_clients.add(ws)
                
                # Start sending updates after subscribing
                await send_updates(ws)

        except Exception as e:
            print(f"Error: {e}")
            await ws.close()

async def send_updates(ws):
    """Send periodic updates to the client after subscribing"""
    while ws in connected_clients:
        try:
            # Send order book update every 5 seconds
            await asyncio.sleep(5)
            await ws.send(json.dumps(order_book_update))
        except websockets.ConnectionClosed:
            print("Client disconnected")
            connected_clients.remove(ws)
            break

async def start_server():
    """Start WebSocket server"""
    server = await websockets.serve(subscribe, "localhost", 9090)
    print("WebSocket server started on ws://localhost:9090")
    await server.wait_closed()

if __name__ == "__main__":
    asyncio.run(start_server())
