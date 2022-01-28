const { generateProducts } = require("./generators/generateProducts");
const { generateOrder } = require("./generators/generateOrder");

const { createProduct, refillStock } = require("./apis/productsApi");
const { createOrder } = require("./apis/ordersApi");

async function mockClients() {
  const products = generateProducts();

  console.log("Creating products");
  products.forEach(async (product) => {
    await createProduct(product);
  });

  setInterval(() => {
    console.log("Refilling stock...");
    products.forEach(async (product) => await refillStock(product));
  }, 12_000);

  setInterval(async () => {
    console.log("Creating order...");
    const order = generateOrder(products);
    await createOrder(order);
  }, 30_000);
}

mockClients().catch((error) => console.log(error));
