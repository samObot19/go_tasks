## **Task Management REST API Documentation**

**overview**

*This API provides a simple task management system with basic CRUD (Create, Read, Update, Delete) operations. It allows users to manage tasks by creating, retrieving, updating, and deleting them.*


**Endpoints***

### Get All Tasks

**Endpoint:** `GET /tasks`
**Description:** *Retrieves a list of all tasks.*
**Response:**

```json
{[
  {
    "id": "1",
    "title": "Finish report",
    "description": "Complete the quarterly report",
    "dueDate": "2023-08-15",
    "status": "in progress"
  },
  {
    "id": "2",
    "title": "Grocery shopping",
    "description": "Buy milk, eggs, and bread",
    "dueDate": "2023-08-12",
    "status": "pending"
  }
]}
```

### Get Task by ID

**Endpoint:** `GET /tasks/{id}`  
**Description:** Retrieves the details of a specific task.
**Response:**

```json
{
  "id": "1",
  "title": "Finish report",
  "description": "Complete the quarterly report",
  "dueDate": "2023-08-15",
  "status": "in progress"
}
```
### Update a Task

**Endpoint:** PUT /tasks/{id}

**Description:** Updates an existing task.

**Request Body:**

```json
{
  "title": "Attend weekly meeting",
  "description": "Weekly team meeting to discuss project progress",
  "dueDate": "2023-08-22",
  "status": "in progress"
}
```
**Response:**

```json
{
  "id": "3",
  "title": "Attend weekly meeting",
  "description": "Weekly team meeting to discuss project progress",
  "dueDate": "2023-08-22",
  "status": "in progress"
}
```
**Delete a Task**

**Endpoint:** DELETE /tasks/{id}

**Description:** Deletes a specific task.

**Response:** 204 No Content


## *Running the API*

**1.** *Ensure you have Go and the Gin Framework installed.*

**2.** *Navigate to the project directory.*

**3.** *Run the following command to start the API server.*


``` bash
go run main.go
```

*The API will be available at http://localhost:8080.*


## *Testing the API*

*You can use tools like Postman, cURL, or any other HTTP client to test the API endpoints.*

## *Documentation*

*The API documentation is available in the docs/documentation.md file.*

## *Future Enhancements*

*Add authentication and authorization features.*

*Implement more advanced task management features (e.g., subtasks, comments, attachments).*

*Improve error handling and logging.*

*Api documentation [here](https://documenter.getpostman.com/view/31283115/2sA3rwMZof)



