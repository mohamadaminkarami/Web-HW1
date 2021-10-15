const express = require("express");
const { PORT } = require("./config");

const redis = require("./redis");

const app = express();
app.use(express.json());

app.get("/node/sha256", async (req, res) => {
  const { encoded } = req.query;

  const rawString = await redis.get(encoded);

  if (rawString) {
    res.send({ raw_string: rawString });
    return;
  }
  res.statusCode = 404;
  res.send({ errors: ["sha256 hash not found!"] });
});

app.listen(PORT, () => {
  console.log(`codec app started on port ${PORT}`);
});
