const { productsClient } = require("./clients");

async function createProduct(product) {
  try {
    await productsClient.post("/", product);
  } catch (error) {}
}

async function refillStock(product) {
  const value = Math.floor(Math.random() * 100);
  try {
    await productsClient.patch(`/${product.id}/stock/${value}`);
  } catch (error) {}
}

module.exports = {
  createProduct,
  refillStock,
};
