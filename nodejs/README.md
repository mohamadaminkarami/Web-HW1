## Introduction
  Nodejs web server app to encode and decode strings to sha256 hashes.

## Installation

After cloning project, you need to add .env file next to package.json:
```
// .env file
PORT= // example:3000
REDIS_HOST= // example:127.0.0.1
REDIS_PORT= // example:6379
REDIS_PASSWORD= // example:password

```

## Running project(Prodution mode)

After cloning project, Run following commands:

```
// using npm
$ npm install
$ npm start

// using yarn
$ yarn install
$ yarn start
```

## Running project(Development mode)

After cloning project, Run following commands:

```
// using npm
$ npm install
$ npm dev

// using yarn
$ yarn install
$ yarn dev
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
Otherwise you will get encoded hash in hex:
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


