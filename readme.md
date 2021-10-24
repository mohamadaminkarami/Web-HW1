Web Programming Fall 2021, 1st Homework.

## Installation

After cloning project, you may add a .env file next to docker-compose.yml:
```
// .env file
NODE_REPLICAS=1
NODE_CPUS=0.125 
NODE_MEM=100M
NODE_CPUS_RES=0.125
NODE_MEM_RES=50M

GO_REPLICAS=1
GO_CPUS=0.125 
GO_MEM=100M
GO_CPUS_RES=0.125
GO_MEM_RES=50M

REDIS_CPUS=0.5
REDIS_MEM=2G
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