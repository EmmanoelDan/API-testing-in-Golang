# Simple Go API

This is a simple API developed in Go using the Gorilla Mux framework.

## Prerequisites

Make sure you have Go installed on your machine. You can download it from: https://golang.org/dl/

## Installation

1. Clone this repository:

   ```shell
   git clone <REPOSITORY_URL>

2. Navigate to the project directory:

    ```shell
    cd directory_name

3. Download the necessary dependencies:

    ```shell
    go mod download

## Running the API

1. In the project directory, run the following command:

        ```shell
        go run main.go

This will start the API and it will be ready to receive requests.

2. Access the API through the following URL:

        ```shell
        http://localhost:3000


## Endpoints
    - GET /books: Returns all books.
    - POST /books: Creates a new book. The book data should be provided in the request body in JSON format.

## Usage Examples
1. Getting all books:

        ```shell
        curl -X GET http://localhost:3000/books

2. Creating a new book:

        ```shell
        curl -X POST -H "Content-Type: application/json" -d '{"title":"New Book", "author":"Author", "year":2023}' http://localhost:3000/books



