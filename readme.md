Web Programming Fall 2021, 1st Homework.

## Installation

After cloning project, you may add a .env file next to docker-compose.yml:
```
// .env file
GO_REPLICAS=1
NODE_REPLICAS=1
```

## Running project

Run following commands to deploy the project:

```
docker-compose -f docker-compose.yml --env-file .env up --build (-d)
```

to use default options:
```
docker-compose up --build (-d)
```