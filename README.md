# go-restful-example

## Getting Started
Navigate to the project directory

- exec `docker build -t go-restful-example .`
- exec `docker run -p 3000:3000 -tid go-restful-example`

## Running Tests
**Note:** running tests will delete all records from the datebase

With docker running:
- exec `go test ./...`

## Endpoints
- `GET /books` returns a JSON array of all the books

- `POST /books data: JSON book schema` inserts a book with the given data

- `DELETE /books` deletes all the books

- `GET /books/{id}` returns a JSON object containing the book with the given id

- `DELETE /books/{id}` deletes the book with the given id

- `PUT /books/{id} data: JSON book schema` updates the book with the given id with the given book data

- `PUT /book/{id}/checkout` updates the status of the book with the given ID to false

- `PUT /book/{id}/checkin` updates the status of the book with given ID to true

## Things we could improve on
- This application is currently lacking authentication.
- Our list-all function should paginate instead of returning one large response.
- Our api does not enforce any sort of rate limiting.
- At the moment, we do not enforce any strict request data from the POST endpoint for updating. An improvement here would be to repond with an error if the decoded request body does not exactly match our interfaces.
- Our api package has defined CRUD methods that are only compatible with one database table. To scale this app in the future, we may want to abstract our api model to support more objects.
- We're lacking an app configuration, a more permanent solution would be preferred to scale the app. This could impact things like db connection strings and overall reduce hardcoded configuration-based literals
- We're using a sqlite database for storage. We might want a more permanent solution for scale. For example, we could containerize postgresql with docker-compose
  - this would also provide an opportunity to write meaningful unit tests for the database
- There's quite a bit of code redundancy in our api-handler code and test code
- Our error messages could be much more helpful. For example, sending a POST request with no data simply returns "EOF"
- Our unit tests could be expanded to test more functionality rather than just status codes
- Docker file could be updated to fetch from remote repo

## Notes

- Took inspiration from:
  - https://github.com/mingrammer/go-todo-rest-api-example
  - https://www.golangprograms.com/golang-restful-api-using-grom-and-gorilla-mux.html
  - https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql
