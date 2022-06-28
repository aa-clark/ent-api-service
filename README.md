# entproject

An API for interacting with user-created arbitrary "services".

## Design

### API Requirements

```
GET    /api/services                         # List services (filter, sort, page, order)
POST   /api/services                         # Create service
GET    /api/services/:id                     # Get service
POST   /api/services/:id/version             # Create a "version" of a service config
PUT    /api/services/:id                     # Update service
DELETE /api/services/:id                     # Delete service
```

### Considerations


Data Model:

I used the [ent](https://entgo.io/) framework to handle modeling and querying data. This was a fun project to explore! sqlite3 was used for simple persistence.

Routing:

I chose Chi, a lightweight router similar to Gorilla/mux. This dependency could absolute be avoided using the standard library and a solution like a regex table to match routes. I opted to write my own JSON marshal functions for server responses instead of Chi's option - and would have greatly liked to implement Protobuf for stricter request/response messages

Improvements:

1. Implement Protocol Buffers for tighter request and response schemas
2. Better separation of duties between server and data layer
3. Add endpoints for the U(pdate) and D(elete) in CRUD
4. Configurable server options (database, port, etc)
5. Middleware: logging, request ids, auth
6. Better testing with cleaner mock data and more coverage

## Running

```
make build run
```

## Test Coverage

```
$ make test_coverage

go test ./... -coverprofile=coverage.out
ok      entproject      0.186s  coverage: 62.2% of statements
?       entproject/database     [no test files]
```

List  services

```
❯ curl -s http://localhost:8080/services\?page\=0\&count\=2\&sort_by\=id\&order_by\=asc
{
    "services": [
        {
            "id": "ad90b71d-cc89-48c0-be3e-52c481576b0e",
            "name": "Prometheus",
            "description": "",
            "versions": [1, 2]
        },
        {
            "id": "f418c4fa-6a1f-478b-a270-9babbfc0048b",
            "name": "Grafana",
            "description": "",
            "versions": [1, 2]
        }
    ],
    "count": 3,
    "prev": "",
    "next": "1"
}
```

Create service

```
❯ curl -s -X POST 'http://localhost:8080/services' \
--data-raw '{
        "name": "Robots first Service",
        "description": "Very apt description"
}'
{
    "id": "59da87a0-ed7e-4d08-a4de-a829653fae05"
}
```

Get Service

```
❯ curl -s 'http://localhost:8080/services/5ef1a6c8-37aa-4c09-8939-ef9ef1850043'
{
    "id": "5ef1a6c8-37aa-4c09-8939-ef9ef1850043",
    "name": "Robots first Service",
    "description": "Very apt description",
    "versions": []
}
```
