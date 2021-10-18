const express = require("express");
const { createHash } = require("crypto");
const { PORT } = require("./config");

const redis = require("./redis");

const app = express();
app.use(express.json());

app.get("/sha", async (req, res) => {
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

app.post("/sha", async (req, res) => {
  const { raw_string: rawString } = req.body;

  if (rawString.length >= 8) {
    const hash = createHash("sha256");

    const encoded = hash.update(rawString).digest("hex");

    await redis.set(encoded, rawString);
    res.send({ encoded });
    return;
  }

  res.statusCode = 400;
  res.send({ errors: ["raw_staring must be at least 8 characters"] });
});