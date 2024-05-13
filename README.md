# Books REST-API

This repository is designed as a learning project to develop a basic RESTful API for managing a collection of books. The API will allow users to add and read the stored book records. operations on book records are stored in the server.

From: [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)

## Features

- Create books by sending POST request on `localhost:8080/books`
- Retrieve books by sending GET request on `localhost:8080/books`
- Retrieve specific book by sending GET request on `localhost:8080/books/:id`, replace `:id` with the book ID.

## Examples

- send `POST` request on `localhost:8080/books`
- send `GET` request on `localhost:8080/books` to get all books.
- send `GET` request on `localhost:8080/books/1` to get book with ID `1`.

## Technologies Used

- Go Programming Language.
- Gin as HTTP web framework for Go.

## Setup

- Clone the repository.
- Open the repository folder.
- Run the server by `go run .` command.
- Send request to the server via `GET` or `POST` request on `localhost:8080/books`
