# To-Do App

This is a simple To-Do application consisting of a React frontend and a Go backend, with MongoDB as the database.

## Project Structure
```
/todo-app
├── client # Frontend code (React)
├── server # Backend code (Go)
├── init-mongo.js # MongoDB initialization script
├── docker-compose.yml
└── README.md
```

## Prerequisites

- Docker
- Docker Compose

## Environment Variables

Create a `.env` file in the project root directory with the following content:

```env
DB_URI=mongodb://root:root1234@mongodb:27017/test?authSource=admin
DB_NAME=test
DB_COLLECTION_NAME=todolist
```

## Building and Running with Docker Compose
To build and run the application with Docker Compose, follow these steps:

1. Build and start the services:
```
    docker-compose up --build
```

2. Build and start the services:
```
    docker-compose up
```

You should see output indicating that both the db and app services are running.

3. Access the application:
```
    Open your web browser and navigate to http://localhost:3000.
```

Services
    - MongoDB: Runs on port 27017
    - Backend (Go): Runs on port 9000
    - Frontend (React): Runs on port 3000

## API Endpoints
```
    GET /api/task: Retrieve all tasks
    POST /api/task: Create a new task
    PUT /api/task/{id}: Mark a task as complete
    PUT /api/undoTask/{id}: Undo a task
    DELETE /api/deleteTask/{id}: Delete a task
    DELETE /api/deleteAllTasks: Delete all tasks
```

## Troubleshooting
- Check Docker Compose Logs:
```
    docker-compose logs app
```

- Rebuild Docker Images:
```
    docker-compose down
    docker-compose build --no-cache
    docker-compose up