# Instruction

1. How to prepare the environment, e.g. Go version, the virtual environment necessary (Docker/Vagrant), etc.

   - Go Version : go version go1.18 windows/amd64
   - Docker Dekstop Version : docker dekstop 4.8.1

2. How to run the storage system, e.g. docker-compose -f docker-storage.yaml up -d

   - ``docker-compose up -d`` (if you want to run it via docker)

3. How to run the application, e.g. go run main.go or make run

   - ``go run main.go`` (if you want to run it in local -> using port 8080)

4. Addresses of the API, e.g. localhost:8080/v1/farm, localhost:8080/v1/pond

   - Create Farm (``localhost:8080/v1/farms`` - POST)
```json
{
    "title": "farm 2",
    "description": "farm catfish"
}
```

   - Get Detail Farm (``localhost:8080/v1/farms/:id`` - GET)
   - Get List Farm (``localhost:8080/v1/farms`` - GET)
   - Update Farm (``localhost:8080/v1/farms/:id`` - PUT)
```json
{
    "title": "farm 1",
    "description": "test update"
}
```

   - Delete Farm (``localhost:8080/v1/farms/:id`` - DELETE)
   - Create Pond (``localhost:8080/v1/ponds`` - POST)
```json
{
    "title": "pond 1",
    "farm_id": 6,
    "description": "pond catfish"
}
```

   - Get Detail Pond (``localhost:8080/v1/ponds/:id`` - GET)
   - Get List Pond (``localhost:8080/v1/ponds`` - GET)
   - Update Pond (``localhost:8080/v1/ponds/:id`` - PUT)
```json
{
    "title": "pond 2",
    "description": "test update"
}
```

   - Delete Pond (``localhost:8080/v1/ponds/:id`` - DELETE)
   - Counter Each Endpoint (``localhost:8080/counter`` - GET)


`5. Any additional information that you think we should know before running the application`

- How to run all unit test:  ``go test ./... -cover``
- This API runs MySQL database from localhost, so prepare the database in your localhost first (you can use delos.sql in this folder)
- Please create Farm data first before creating data on Pond, because if it is not created first it will cause an error
- if you want to use POSTMAN Collection, it's in this folder (Test Delos.postman_collection.json)
