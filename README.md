# rest-api-go
for learning purposes

## Create User
```go
curl --location --request POST 'http://localhost:6001/users' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'name=Yuda'
```

## Read Users
```go
curl --location --request GET 'http://localhost:6001/users'
```

## Update User
```go
curl --location --request PUT 'http://localhost:6001/users' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'id=1' \
--data-urlencode 'name=Fajar'
```

## Delete User
```go
curl --location --request DELETE 'http://localhost:6001/users?id=1'
```