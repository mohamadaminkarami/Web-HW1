const express = require("express");
const { createHash } = require("crypto");
const { PORT } = require("./config");

const redis = require("./utils/redis");
const validateEncodedString = require("./validatiors/encoded-string-validator");
const validateRawString = require("./validatiors/raw-string-validator");

const app = express();
const router = express.Router();
router.use(express.json());

router.get("/sha", async (req, res) => {
  const { encoded } = req.query;

  try {
    validateEncodedString(encoded);
  } catch (err) {
    res.statusCode = 400;
    res.send({ errors: [err.message] });
  }

  const rawString = await redis.get(encoded);

  if (rawString) {
    res.send({ raw_string: rawString });
    return;
  }
  res.statusCode = 404;
  res.send({ errors: ["sha256 hash not found!"] });
});

router.post("/sha", async (req, res) => {
  const { raw_string: rawString } = req.body;

  try {
    validateRawString(rawString);
  } catch (err) {
    console.log(err)
    res.statusCode = 400;
    res.send({ errors: [err.message] });
    return;
  }

  const hash = createHash("sha256");

  const encoded = hash.update(rawString).digest("hex");

  await redis.set(encoded, rawString);
  res.send({ encoded });
});

app.use(router);
app.listen(PORT, () => {
  console.log(`codec app started on port ${PORT}`);
});
app.use((err) => {
  console.log(err);
});
process.on("error", (err) => {
  console.log(err);
});
