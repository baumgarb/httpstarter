# Get all todos
curl http://localhost:8080/api/v1/todos | jq

# Post new todo
curl http://localhost:8080/api/v1/todos --data '{"name": "do something else"}'

# Update existing todo
curl http://localhost:8080/api/v1/todos/3 --data '{"id": 3, "name": "Play Tennis - Changed", "done": false}' -X PUT