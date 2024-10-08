# Golculator API

**Golculator API** is a simple and efficient RESTful API for performing basic arithmetic operations such as addition, subtraction, multiplication, and division. Built with Go (Golang), this API provides an easy-to-use interface for developers to integrate math functionalities into their applications.

## Features

- Perform basic arithmetic operations:
  - Addition
  - Subtraction
  - Multiplication
  - Division
- Handle errors gracefully with appropriate status codes and messages
- Input Validation Ensures query parameters are present and valid, handling errors such as invalid numbers and division by zero gracefully.
- Modular Tests: Each module (math operations, API utilities, handlers) is thoroughly tested with unit tests, ensuring functionality and reliability across the application.
- Return results in JSON format

## Getting Started

### Prerequisites

- Go (version 1.16 or later)
- A web browser or an API client like Postman or curl

### Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/mahmoudalnkeeb/github.com/mahmoudalnkeeb/golculator-api.git
   cd github.com/mahmoudalnkeeb/golculator-api
   ```

2. **Install dependencies** (if any):

   ```bash
   go mod tidy
   ```

3. **Run the API**:

   ```bash
   go run .
   ```

   The API will be available at `http://localhost:3333`.

### API Endpoints

#### Addition

- **Endpoint**: `/add`
- **Method**: `GET`
- **Query Parameters**:
  - `n1`: First number (float)
  - `n2`: Second number (float)
- **Example**: `/add?n1=10&n2=5`
- **Response**:
  - `200 OK`: JSON response with result

    ```json
    {
      "result": 15
    }
    ```

#### Subtraction

- **Endpoint**: `/subtract`
- **Method**: `GET`
- **Query Parameters**:
  - `n1`: First number (float)
  - `n2`: Second number (float)
- **Example**: `/subtract?n1=10&n2=5`
- **Response**:
  - `200 OK`: JSON response with result

    ```json
    {
      "result": 5
    }
    ```

#### Multiplication

- **Endpoint**: `/multiply`
- **Method**: `GET`
- **Query Parameters**:
  - `n1`: First number (float)
  - `n2`: Second number (float)
- **Example**: `/multiply?n1=10&n2=5`
- **Response**:
  - `200 OK`: JSON response with result

    ```json
    {
      "result": 50
    }
    ```

#### Division

- **Endpoint**: `/divide`
- **Method**: `GET`
- **Query Parameters**:
  - `n1`: First number (float)
  - `n2`: Second number (float)
- **Example**: `/divide?n1=10&n2=5`
- **Response**:
  - `200 OK`: JSON response with result

    ```json
    {
      "result": 2
    }
    ```

  - `400 Bad Request`: If `n2` is `0`

    ```text
    Division by zero
    ```

### Error Handling

The API provides meaningful error messages for invalid requests:

- `400 Bad Request`: For invalid query parameters or division by zero.

### Running Tests

To run the tests for the API, use the following command:

```bash
go test ./tests
```

### License

This project is licensed under the MIT License - see the [MIT](LICENSE) file for details.
