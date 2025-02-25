# Real-Time Notification System

## Overview

A real-time notification system built using Kafka, Redis, PostgreSQL, and WebSockets. This backend service ensures instant event-driven notifications with efficient storage and retrieval.

## Features

- **Kafka Producer & Consumer**: Events are produced and consumed asynchronously.
- **PostgreSQL Storage**: Notifications are stored for persistence.
- **Redis Caching**: Unread notifications are cached for quick access.
- **WebSockets**: Real-time push notifications to users.
- **REST API**: Fetch notifications via user ID.

## Tech Stack

- **Go**: Backend service.
- **Kafka**: Message broker for event-driven processing.
- **PostgreSQL**: Persistent storage.
- **Redis**: Fast in-memory cache.
- **WebSockets**: Real-time communication.


## Architecture Flow

1️⃣ **Notification Event Ingestion**  
   - The API receives notification data and **publishes it to Kafka** for reliable message handling.

2️⃣ **Event Processing & Storage**  
   - A Kafka Consumer **processes the event**, storing full notification details in **PostgreSQL**.  
   - The notification **ID is cached in Redis** for quick access to unread notifications.

3️⃣ **Fetching Notifications**  
   - When users fetch notifications, the API retrieves **IDs from Redis** first (low-latency response).  
   - If needed, full details are fetched from **PostgreSQL**.

4️⃣ **Real-time Notification Delivery**  
   - WebSockets **push real-time notifications** to actively connected clients.  
   - If a user **views a notification live**, it is marked as **read** immediately in Redis and PostgreSQL.

5️⃣ **Manual Read Acknowledgment**  
   - Users can manually mark notifications as **read** via an API.  
   - The system **removes the notification from Redis** and updates PostgreSQL.



## Installation

1. Clone the repo:
   ```sh
   git clone https://github.com/likhithkp/realtime-notifications-service.git
   ```
2. Set up Kafka, Redis, and PostgreSQL.
3. Configure `.env` with necessary environment variables.
4. Run the service:
   ```sh
   go run main.go
   ```

## API Endpoints

- `POST /createUser` - Create a user first.
- `POST /createNotification` - Publish a notification.
- `GET /notifications/{user_id}` - Fetch notifications for a user.
- `GET /ws/{user_id}` - Listen to live notifications.


## License

MIT License.
