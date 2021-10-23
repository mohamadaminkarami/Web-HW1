## Introduction
  Golang web server app to encode and decode strings to sha256 hashes.

## Installation

After cloning project, you need to add .env file next to main.go:
```
// .env file
SERVER_ADDR= // default:localhost:8080
REDIS_ADDR= // default:localhost:6379
REDIS_PASSWORD= // default:no password set
REDIS_DB= // default:0
```

## Running project

Run following commands:

```
$ go mod tidy
$ go build main.go
$ go run main.go
```

## Routes
It provides following routes:

### POST /sha

You can encode your strings(at least 8 characters) to sha256 hash with.
Send POST request with body:
```json
{
    "raw_string": "at-least-8-character string"
}
```
If your string is less than 8 characters web server returns following response:
```json
{
  "errors": ["raw_staring must be at least 8 characters"]
}
```
Otherwise, you will get encoded hash in hex:
```json
{
    "encoded": "encoded hash in hex"
} 
```

### GET /sha

You can decode your sha256 hashes with.
Send GET request with query params:
```js
/sha?encoded=${yourhash}
```
If your hash does not exist on database response will be:
```json
{
  "errors": ["sha256 hash not found!"]
}
```
Otherwise you will get decoded string:
```json
{
    "raw_string": "decoded string"
} 
```


