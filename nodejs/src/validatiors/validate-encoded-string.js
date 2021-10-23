function validateEncodedString(encoded) {
  if (!encoded) {
    throw new Error("encoded param is required");
  }
  if (typeof encoded !== "string") {
    throw new Error("encoded param must be String");
  }
  if (encoded.length === 0) {
    throw new Error("encoded param cannot be 0-length-string");
  }
}

module.exports = validateEncodedString;
