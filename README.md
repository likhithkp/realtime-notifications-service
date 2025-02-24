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
1. API receives notification data and publishes it to Kafka.
2. Kafka Consumer processes the event, storing details in PostgreSQL and caching notification IDs in Redis.
3. Users fetch notifications via API, retrieving IDs from Redis and full details from PostgreSQL if needed.
4. WebSockets push real-time notifications to connected clients.

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
- `POST /notifications` - Publish a notification.
- `GET /notifications/{user_id}` - Fetch notifications for a user.

## License
MIT License.
