# Basic CRUD API Golang

This is a basic CRUD API using Go. It uses an in-memory slice for storing movies and allows you to perform CRUD operations on them.

## Running locally

```bash
go mod download # install dependencies
go run main.go # run it!
```

## Local testing

### `GET /movies`

```
curl --request GET \
  --url http://localhost:8080/movies
```

### `GET /movie`

```
curl --request GET \
  --url http://localhost:8080/movie/1
```

### `PUT /movie/{id}`

```
curl --request PUT \
  --url http://localhost:8080/movie/1 \
  --header 'Content-Type: application/json' \
  --data '{
	"id": "1",
	"isbn": "12345",
	"title": "My first movie - updated!",
	"director": {
		"firstName": "John",
		"lastName": "Doe"
	}
}'
```

### `POST /movie`

```
curl --request POST \
  --url http://localhost:8080/movie \
  --header 'Content-Type: application/json' \
  --data '{
	"isbn": "23456",
	"title": "a third movie",
	"director": {
		"firstName": "John",
		"lastName": "Doe"
	}
}'
```

### `DELETE /movie/{id}`

```
curl --request DELETE \
  --url http://localhost:8080/movie/1
```

