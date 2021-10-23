Web Programming Fall 2021, 1st Homework.

## Installation

After cloning project, you need to add .env file next to docker-compose.yml:
```
// .env file
GO_REPLICAS=1
NODE_REPLICAS=1

```

## Running project

After cloning project, Run following commands:

```
docker-compose -f docker-compose.yml --env-file .env up --build
```