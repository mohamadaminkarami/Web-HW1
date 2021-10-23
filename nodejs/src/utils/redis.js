const Redis = require("ioredis");
const { REDIS } = require("../config");

const redis = new Redis(REDIS);

redis
  .ping()
  .then((res) =>
    res === "PONG"
      ? console.log(
          `redis connected successfully on host=${REDIS.host} port=${REDIS.port}`
        )
      : console.log("Cannot connect to redis")
  );

module.exports = redis;
