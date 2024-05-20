# bitgo


## Curl for APIs

### Create Notification
```
curl --location 'localhost:8081' \
--header 'Content-Type: application/json' \
--data '{

    "CurrentPrice": 1344.0,
    "Status": 0
}'
```


### Create Notification
send status empty or not to fetch all notifications
```
curl --location 'localhost:8081?status=Failed' \
--data ''
```

### Delete Notification
```
curl --location --request DELETE 'localhost:8081/0b27dbbd-bede-4c45-8605-3f45e2d40ed8' \
--data ''
```

## TODOs:

1. worker status update in db post sending mail
2. Worker can be observer pattern based on some param
3. Worker main.go can be another server independent


## Steps to run:

```azure
go run cmd/server/main.go
```
