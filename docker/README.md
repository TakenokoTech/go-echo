# Docker

```shell
docker-compose up -d
```

```shell
$ curl -s -X GET "http://localhost:8080/v1/users" -H "Content-Type: application/json"
$ curl -s -X POST "http://localhost:8080/v1/users" -H "Content-Type: application/json" -d '{"UserName": "b"}'
$ curl -s -X GET "http://localhost:8080/v1/users/1" -H "Content-Type: application/json"
```
