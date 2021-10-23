function validateRawString(rawString) {
  if (!rawString) {
    throw new Error("raw_string must be provided!");
  }
  if (typeof rawString !== "string") {
    throw new Error("raw_string must be String!");
  }
  if (rawString.length < 8) {
    throw new Error("raw_staring must be at least 8 characters");
  }
}

module.exports = validateRawString;
