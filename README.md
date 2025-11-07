Key-Value Persistent Store (Golang)
Overview
This is a simple key-value store built in Go.
It provides an HTTP API to save and retrieve data, with persistence to a file (data.json) so that all data remains even after server restarts.
How to Run
1. Without Docker
1. Install Go.
2. Open terminal/PowerShell and navigate to your project folder:
   cd keyvalue-store
3. Run the program:
   go run main.go
4. If you see:
    Server started on port 8080
   it means your server is running successfully.
2. With Docker (Optional)
1. Make sure you have Docker and Docker Compose installed.
2. In the project folder, run:
   docker compose up --build
3. You should see the server starting logs in Docker console.
4. To stop the server, press CTRL+C and then run:
   docker compose down
API Endpoints
1. PUT /objects
Save data with a key-value pair.
Example:
curl -X PUT http://localhost:8080/objects \
-H "Content-Type: application/json" \
-d '{"key": "user:1234", "value": {"name": "Amin Alavi", "age": 23, "email": "a.alavi@fum.ir"}}'
Response: HTTP 200 OK
2. GET /objects/{key}
Retrieve data by key.
Example:
curl http://localhost:8080/objects/user:1234
Response:
{"name":"Amin Alavi","age":23,"email":"a.alavi@fum.ir"}
Persistence
All data is saved in data.json to ensure it persists even after restarting the server or the Docker container.
Project Files
Make sure your project folder contains:
- main.go
- data.json
- Dockerfile
- docker-compose.yml
- README.docx
Quick Commands Summary
Without Docker
cd keyvalue-store
go run main.go
With Docker
docker compose up --build
docker compose down
Test API
# Save data
curl -X PUT http://localhost:8080/objects \
-H "Content-Type: application/json" \
-d '{"key": "user:1234", "value": {"name": "Amin Alavi", "age": 23, "email": "a.alavi@fum.ir"}}'

# Retrieve data
curl http://localhost:8080/objects/user:1234
# keyvalue-store
