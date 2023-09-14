# Person API Documentation

This documentation provides information on how to use the Person API to perform CRUD (Create, Read, Update, Delete) operations on person records.

## Table of Contents
- [API Endpoints](#api-endpoints)
    - [Create a Person](#create-a-person)
    - [Get Person by ID](#get-person-by-id)
    - [Update Person by ID](#update-person-by-id)
    - [Delete Person by ID](#delete-person-by-id)
- [Request and Response Formats](#request-and-response-formats)
    - [Create Person Request](#create-person-request)
    - [Create Person Response](#create-person-response)
    - [Get Person Response](#get-person-response)
    - [Update Person Request](#update-person-request)
    - [Update Person Response](#update-person-response)
- [Usage Examples](#usage-examples)
- [Limitations](#limitations)
- [Setup and Deployment](#setup-and-deployment)

## API Endpoints

### Create a Person

- **Endpoint**: `/api`
- **Method**: POST
- **Description**: Create a new person record.

#### Request

##### Create Person Request

- **HTTP Method**: POST
- **URL**: `http://localhost:8000/api/name`
- **Headers**:
    - Content-Type: application/json
- **Request Body** (JSON):

```json
{
    "name": "John",
    "age": 30,
    "email": "johndoe@example.com"
}
```
#### Response
#### Create Person Response
- Status Code: 200 OK
  - Response Body (JSON):
    
      ```json
      {
          "id": 1,
          "name": "John",
          "age": 30,
          "email": "johndoe@gmail"    
        }
        ```
### Get Person by ID

- Endpoint: /api/{id}
- Method: GET
- Description: Fetch details of a person by their ID.
- Request
- HTTP Method: GET
- URL: http://localhost:8000/api/1
- Response
- Get Person Response
- Status Code: 200 OK
- Response Body (JSON):

```json
{
    "id": 1,
    "name": "John Doe",
    "age": 30,
    "email": "johndoe@gmail.com"  
} 
```

- Update Person by ID
- Endpoint: /api/{id}
- Method: PUT
- Description: Modify details of an existing person by their ID.
- Request
- Update Person Request
- HTTP Method: PUT
- URL: http://localhost:8000/api/1
- Headers:
- Content-Type: application/json
- Request Body (JSON):
- json

 ```json
{
    "name": "Updated Name",
    "age": 35
   
}
``` 


- Response

```json
{
    "id": 1,
    "name": "Updated Name",
    "age": 35,
    "email": "john@gmail.com"

}
```

- Delete Person by ID
- Delete Person by ID
- Endpoint: /api/{id}
- Method: DELETE
 -  Description: Remove a person by their ID.
 -  Request
 -  HTTP Method: DELETE
  - URL: http://localhost:8000/api/1
  - Response
 -  Status Code: 200 OK
 -  Response Body (JSON):
  - json
   
````
  {
  "message": "Person deleted successfully"
  }

````

## Limitations

- The API assumes that each person has a unique ID for fetching, updating, and deleting. Ensure that you provide the correct ID when performing these operations.
- Validation is performed to ensure that certain fields are of the correct data type (e.g., strings). Make sure to provide valid data in your requests.

## Setup and Deployment

1. Clone the repository from GitHub.
2. Install the required dependencies.
3. Configure your database connection settings in the `database/db.go` file.
4. Run the application using `go run main.go`.
5. The API will be accessible at `http://localhost:8000`.





## UML Diagram
![Person Class Diagram](path/to/your/diagram.png)

