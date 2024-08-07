**Task Management API Documentation**

**Base URL**

```bash
http://localhost:8080/api
```

**Authentication and Authorization**

**Authentication**

This API uses JWT (JSON Web Tokens) for authentication. Users must log in to obtain a token and include this token in the Authorization header of each request.

**Authorization**


The API has role-based access control (RBAC). Only users with the admin role can create, update, delete tasks, and promote users.

**Endpoints**

Authentication

**Login**
```bash
post-> api/login
```

**Request Body**

```json
{
  "username": "johndoe",
  "password": "securepassword123"
}
```

**Response Body**

```json
{
  "message": "User logged in successfully",
  "token": "jwt-token"
}
```

**Tasks**

**Create Task (Admin Only)**
```bash
post api/admin/tasks
```
**Headers:**
```makefile
Authorization: Bearer jwt-token
```

**Request Body:**

```json
{
  "title": "Complete Project Report",
  "description": "Finish the final report for the project and submit it by the end of the week.",
  "status": "Pending"
}
```

**Response:**

```json
{
  "message": "Task created"
}
```
**Retrieve Tasks(for all user)**

```bash
GET /tasks
```

**Headers:**
```makefile
Authorization: Bearer jwt-token
```

**Response:**

```json
{
  "tasks": [
    {
      "id": "66b35e8269dbc16fed2ef33b",
      "Title": "Complete Project Report",
      "Description": "Finish the final report for the project and submit it by the end of the week.",
      "DueDate": "0001-01-01T00:00:00Z",
      "Status": "Pending"
    }
  ]
}
```

**Delete Task (Admin Only)**
```bash
DELETE api/admin/tasks/{id}
```

**Headers:**
```makefile
Authorization: Bearer jwt-token
```

**Request Example:**

```bash
localhost:8080/api/admin/tasks/66b35e8269dbc16fed2ef33b
```
**Response**

```json
{
  "message": "Task removed"
}
```

**User Management (Admin Only)**

**Promote User to Admin**

```bash
POST /admin/users/promote/{username}
```

**Headers:**

```makefile
Authorization: Bearer jwt-token
```
**Request Example:**

```bash
localhost:8080/api/admin/users/promote/joshua
```

**Response**
```json
{
  "message": "the user joshua is promoted as an admin successfully!"
}
```

**Update Task by ID (Admin Only)**
```bash
PUT /admin/tasks/:id
```

**Headers:**

```makefile
Authorization: Bearer jwt-token
```

**Request Body**

```json
{
  "title": "Complete Project Report",
  "description": "Finish the final report for the project and submit it by the end of the week.",
  "status": "completed"
}
```
**Response:**

```json
{
  "message": "Task updated"
}
```

**Error handling**

The API returns standard HTTP status codes to indicate the success or failure of an API request. Here are some common status codes:

+ **200 OK:** The request was successful.

+ **201 Created:** The resource was successfully created.
+ **400 Bad Request:** The request was invalid or cannot be otherwise served.
+ **401 Unauthorized:** Authentication failed or user does not have permissions for the desired action.
+ **403 Forbidden:** The authenticated user does not have access to the requested resource.
+ **404 Not Found**: The requested resource could not be found.
+ **500 Internal Server Error:** An error occurred on the server.

