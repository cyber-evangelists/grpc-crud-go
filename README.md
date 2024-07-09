# gRPC CRUD with Swagger Documentation

## Overview
This project demonstrates a basic gRPC CRUD application with Swagger documentation. It includes both gRPC server and client implementations as well as an HTTP server with Swagger integration for API documentation.

## Project Structure
### Files and Directories

- **main.go**: Entry point for the application.
- **server/**: Contains the server-side code.
  - **server.go**: Initializes and starts the gRPC server.
  - **user-service.go**: Implements the user service with CRUD operations.
  - **handler.go**: Contains request handlers.
- **proto/**: Contains the protobuf definitions.
  - **user.proto**: Protobuf definition file for the user service.
- **docs/**: Contains the Swagger documentation.
  - **json.docs**: Swagger JSON documentation file.
- **client/**: Contains the client-side code.
  - **client.go**: Implements a gRPC client to interact with the server.
 

## Installation

### Prerequisites
- Go 1.16+
- PostgreSQL
- Swagger (for generating documentation)

### Steps
1. Clone the repository:
    ```sh
    git clone https://github.com/your-repo/swag-grpc-crud.git
    cd swag-grpc-crud
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Set up PostgreSQL database:
    - Create a database and update the connection string in `server/server.go`.

4. Generate gRPC code from proto file:
    ```sh
    protoc --go_out=. --go-grpc_out=. proto/user.proto
    ```

5. Generate Swagger documentation:
    ```sh
    swag init
    ```

## Usage

### Start the Servers
Start the gRPC and HTTP servers:
```sh
go run main.go
 ```

## Using the gRPC Client
You can test the gRPC server using the client provided in client/client.go or any gRPC client of your choice.

## Accessing Swagger Documentation
Once the HTTP server is running, access the Swagger UI at http://localhost:8080/swagger/index.html to see the API documentation and test the endpoints.


## Conclusion

This project showcases a complete implementation of a gRPC-based CRUD application, enhanced with Swagger for API documentation. By following the provided structure and setup instructions, you can easily start both gRPC and HTTP servers, connect to a PostgreSQL database, and interact with the API using a gRPC client or Swagger UI. The project structure is modular, making it easy to extend or modify as needed for other types of gRPC services or additional features. This setup ensures that you have a robust and well-documented foundation for building scalable microservices with gRPC and Go.
## Acknowledgements

```bash
This app was made with 💖 by Hamza under the guidance of Sir Husnain.
```
