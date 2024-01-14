## Description

Client-Service is an Application Programming Interface (API) to handle add,update,delete and get client of ZOHS company.
## Installation :
```bash
# install requirements
$ go get -u github.com/go-sql-driver/mysql github.com/jinzhu/inflection github.com/gorilla/mux github.com/joho/godotenv github.com/jinzhu/inflection github.com/dgrijalva/jwt-go
```
## Running the app : 
```bash
# Run application
$ go run main.go
```
## Build Docker image : 
```bash
# build a docker image
$ docker build -t client-service .
```
## Running the app in the Docker : 
```bash
# Run docker image
$ docker run -p 5000:5000 client-service
```
## Running with Docker compose :
```bash
# Run docker compose
$ docker compose up
```

## Models

### Client

- `ID: string` - Unique identifier for the client.
- `FirstName: string` - First name of the client.
- `LastName: string` - Last name of the client.
- `NationalID: string` - National ID of the client.
- `Email: string` - Email address of the client.
- `BirthDate: string` - Birth date of the client.
- `City: string` - City of the client.
- `Nationality: string` - Nationality of the client.
- `Gender: string` - Gender of the client.
- `CreateDate: string` - Date of client creation.
- `LastModificationDate: string` - Date of the last modification.
- `Status: string` - Status of the client.
- `Address: string` - Address of the client.

## Available Endpoints

### 1.GetClients

- **Endpoint:** `/client`
- **Method:** GET
- **Description:** Get all clients.
- **Response:**
  - `200`: Successful retrieval with an array of all clients.
  - `500`: Internal Server Error.

### 2.GetClient

- **Endpoint:** `/client/{id}`
- **Method:** GET
- **Description:** Get a client by ID.
- **Response:**
  - `200`: Successful retrieval with client details.
  - `404`: Client not found.

### 3.CreateClient

- **Endpoint:** `/client`
- **Method:** POST
- **Description:** Create a new client.
- **Request Body:**
  - Provide the client details in the request body.
- **Response:**
  - `201`: Client created successfully with client details.
  - `500`: Internal Server Error.

### 4.UpdateClient

- **Endpoint:** `/client/{id}`
- **Method:** PUT
- **Description:** Update a client by ID.
- **Request Body:**
  - Provide the updated client details in the request body.
- **Response:**
  - `200`: Client updated successfully.
  - `500`: Internal Server Error.

### 5.DeleteClient

- **Endpoint:** `/client/{id}`
- **Method:** DELETE
- **Description:** Delete a client by ID.
- **Response:**
  - `200`: Client deleted successfully.
  - `404`: Client not found.
  - `500`: Internal Server Error.




## Stay in touch :
- Author - [Ouail Laamiri](https://www.linkedin.com/in/ouaillaamiri/) , [Hajar Sadik](https://www.linkedin.com/in/hajar-sadik-b27594268?miniProfileUrn=urn%3Ali%3Afs_miniProfile%3AACoAAEGU9k8BLhHzARArf7SblDplgU6ufFqc-0w&lipi=urn%3Ali%3Apage%3Ad_flagship3_search_srp_all%3BwRckppvzRNiFd9xB2PswKw%3D%3D)
- Documentation - [Postman](https://documenter.getpostman.com/view/27440422/2s9YkuXxNb)

## License

Client-Service is [GPL licensed](LICENSE).


