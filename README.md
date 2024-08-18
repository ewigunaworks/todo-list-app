# Todo List App API
A simple todo list app system API built using Go Language and Echo Framework. This API allows you to Create, Retrieve, Update and Delete todo task. 

## Features
1. Create Todo: Add new todo task to the system.
2. Retrieve Todo: Fetch todo task details by ID.
3. Update Todo: Modify existing todo task details.
4. List Todo: Retrieve a list of all todo tasks with optional filters.
5. Delete Todo : Delete todo task from system
 

## Getting Started
## Prerequisites
1. Go 1.21.x
2. MySQL 5.7

## Installation

1. Clone the repository
2. Create a MySQL database named `todolist`
3. System will automigrate the table, or Run the DDL Query in "query.sql"
3. Update the database connection string in `.env`
4. Install dependencies
    ```sh
    go mod tidy
    ```
5. Run the API
   ```sh
   go run main.go
   ```
   

## Endpoint
### Create Todo
Endpoint: POST /api/v1/todo/create

Description: Create a new todo task.

Request Body: 
```json
{
    "Title": "Test",
    "Description":"Desc",
    "DueDate": "2024-08-22",
    "Status": 1 //0: waiting list, 1: in progress, 2: done
}
```
Response:
```json
{
    "Status": "success",
    "StatusCode": 200,
    "Message": "Success Create Todo List",
    "Data": {
        "id": 2,
        "title": "Test",
        "description": "Desc",
        "status": "Waiting List",
        "dueDate": "2024-08-16 06:30:30",
        "createdAt": "2024-08-18 07:56:55",
        "updatedAt": "",
        "deletedAt": ""
    }
}
```
### Update Todo
Endpoint: PUT  /api/v1/todo/update/{id}

Description: Update an existing todo task.

Request Body: 
```json
{
    "Title": "Test 1 upd",
    "Description":"Desc upd",
    "DueDate": "2024-08-16 06:52:30",
    "Status": 1 //0: waiting list, 1: in progress, 2: done
}
```
Response:
```json
{
    "Status": "success",
    "StatusCode": 200,
    "Message": "Success Update Todo List",
    "Data": {
        "id": 1,
        "title": "Test 1 upd",
        "description": "Desc upd",
        "status": "In Progress",
        "dueDate": "2024-08-16 13:52:30",
        "createdAt": "2024-08-18 07:54:44",
        "updatedAt": "2024-08-18 07:58:06",
        "deletedAt": ""
    }
}
```
### Get Todo
Endpoint: GET /api/v1/todo/{id}

Description: Retrieve todo task by ID.

Request Body: 

Response:
```json
{
    "Status": "success",
    "StatusCode": 200,
    "Message": "Success Get Detail Todo List",
    "Data": {
        "id": 1,
        "title": "Test 1",
        "description": "Desc",
        "status": "Waiting List",
        "dueDate": "2024-08-16 13:30:13",
        "createdAt": "2024-08-18 07:52:07",
        "updatedAt": "",
        "deletedAt": ""
    }
}
```
### Get List Todo
Endpoint: GET /api/v1/todo/list

Description: Retrieve an todo task by its ID.

Query Parameters:
- Search
- FilterField (filter by Title, Status, Description, DueDate, CreatedAt)
- FilterValue
- Page
- Limit

Response:
```json
{
    "Status": "success",
    "StatusCode": 200,
    "Message": "Success Get List Todo List",
    "Data": [
        {
            "id": 4,
            "title": "Test",
            "description": "Desc",
            "status": "In Progress",
            "dueDate": "2024-08-22",
            "createdAt": "2024-08-18 08:22:14",
            "updatedAt": "",
            "deletedAt": ""
        },
        {
            "id": 3,
            "title": "Test",
            "description": "Desc",
            "status": "Waiting List",
            "dueDate": "2024-08-22",
            "createdAt": "2024-08-18 08:16:27",
            "updatedAt": "",
            "deletedAt": ""
        },
        {
            "id": 2,
            "title": "Test 1 upd",
            "description": "Desc upd",
            "status": "In Progress",
            "dueDate": "2024-08-16",
            "createdAt": "2024-08-18 08:16:21",
            "updatedAt": "2024-08-18 08:20:58",
            "deletedAt": ""
        },
        {
            "id": 1,
            "title": "Test 1 upd",
            "description": "Desc upd",
            "status": "Done",
            "dueDate": "2024-08-16",
            "createdAt": "2024-08-18 08:16:16",
            "updatedAt": "2024-08-18 08:20:51",
            "deletedAt": ""
        }
    ],
    "Paginate": {
        "CurrentPage": 1,
        "NextPage": 2,
        "TotalPage": 1,
        "PerPage": 10,
        "TotalRow": 4
    }
}
```

### Delete Todo
Endpoint: DELETE /api/v1/todo/delete/{id}

Description: Delete todo task by ID.

Request Body: 

Response:
```json
{
    "Status": "success",
    "StatusCode": 200,
    "Message": "Success Delete Todo List",
    "Data": null
}
```
## Error Handling
The API returns standard HTTP status codes for errors:

- 400 Bad Request: Invalid request data.
- 404 Not Found: Resource not found.
- 500 Internal Server Error: Server error.
