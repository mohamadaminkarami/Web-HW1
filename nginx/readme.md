## Introduction
  Nginx config for routing and Frontend code to work with backends.

## Installation & Running project

Build and run the docker using following commands in the current directory:
```
docker build -t nginx-image-name .
docker run --name nginx-container-name -d -p 8000:80 nginx-image-name
```
attention: to run this container alone you'll need to comment out dependencies on other containers (to run the whole project use docker-compose)

## Paths
It provides following paths on ip:8000

## location /node/sha256/
proxy pass to node_backend_ip/sha

## location /go/sha256/
proxy pass to go_backend_ip/sha

## location /
The html code will be loaded with a simple UI to code and encode string

