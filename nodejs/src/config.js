require("dotenv").config();

module.exports = {
  PORT: parseInt(process.env.PORT, 10),

  REDIS: {
    host: process.env.REDIS_HOST,
    port: parseInt(process.env.REDIS_PORT, 10),
    password: process.env.REDIS_PASSWORD,
  },
};
