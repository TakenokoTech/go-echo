# go-echo

```shell
go install github.com/golang/mock/mockgen
go get github.com/golang/mock/gomock
```

```shell
mockgen -source=player_api.go -destination=./mock/player_api_mock.go -package=repository
```