# HTTP service for qr code creation

## Running the application

```
docker-compose up --build
```

### Running the application without Docker

```
go run .
```

### Running tests (app must be running)

```
go test
```

## HTTP Methods
```
"GET" / — Checking the server connection

    example: 
        "GET" :8080/
```

```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8080/ping
```
```
"GET" /qr — Get a QR Code
    options:
        data — Data to be encoded into a qr code

    example: 
        "GET" :8080/qr?data=https://github.com/pchchv
```

```
"DELETE" /qr — Delete all QR Codes

    example: 
        "DELETE" :8080/qr
```
