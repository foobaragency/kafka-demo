const { generateIds } = require("./generateIds");
const { productInfo } = require("./productInfo");

function generateProducts() {
  const ids = generateIds();

  const products = ids.map((id) => ({
    id,
    ...productInfo,
  }));

  return products;
}

module.exports = { generateProducts };
