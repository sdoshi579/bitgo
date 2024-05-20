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

1. worker status update in db post sending mail (Done)
2. Worker can be observer pattern based on some param
3. Worker main.go can be another server independent
4. Common handler function which add header
5. Common error interface so that can have http status specific to error
6. Do soft delete instead of deleting from map


## Steps to run:

```azure
go run cmd/server/main.go
```
