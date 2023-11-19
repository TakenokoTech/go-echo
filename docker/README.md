# Docker

```shell
docker-compose up -d
```

### users

```shell
$ curl -s -X GET "http://localhost:8080/v1/users" -H "Content-Type: application/json"
$ curl -s -X POST "http://localhost:8080/v1/users" -H "Content-Type: application/json" -d '{"UserName": "b"}'
$ curl -s -X GET "http://localhost:8080/v1/users/1" -H "Content-Type: application/json"

```

### todos

```shell
curl -s -X GET "http://localhost:8080/v1/todos" -H "Content-Type: application/json"
curl -s -X POST "http://localhost:8080/v1/todos" -H "Content-Type: application/json" -d '{"UserId": "1", "Task": "task", "Status": "0"}'
```
