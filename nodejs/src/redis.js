const Redis = require("ioredis");
const { REDIS } = require("./config");

const redis = new Redis(REDIS);

module.exports = redis;
