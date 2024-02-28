# Task Management API

This project implements a Task Management API using the Gogin Framework with a SQLite database. The API provides endpoints for creating, retrieving, updating, and deleting tasks, as well as listing all tasks.

## Getting Started

### Prerequisites

- Go (Golang) installed on your machine
- Git
- SQLite

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/task-management-api.git
    ```

2. Change into the project directory:

    ```bash
    cd task-management-api
    ```

3. Install dependencies:

    ```bash
    go get -v
    ```

4. Run the application:

    ```bash
    go run main.go
    ```

By default, the application will be accessible at http://localhost:8080.

## API Endpoints

- **Create a new task:**
  - Endpoint: `POST /api/tasks`
  - Request Body: JSON payload containing task details (title, description, due date)
  - Response: Created task with assigned ID if successful, error message if not.

- **Retrieve a task by ID:**
  - Endpoint: `GET /api/tasks/:id`
  - Response: Task details if found, error message if not.

- **Update a task by ID:**
  - Endpoint: `PUT /api/tasks/:id`
  - Request Body: JSON payload containing updated task details (title, description, due date)
  - Response: Updated task details if successful, error message if not.

- **Delete a task by ID:**
  - Endpoint: `DELETE /api/tasks/:id`
  - Response: Success message if deletion
