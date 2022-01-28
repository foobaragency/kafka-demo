const { v4: uuidV4 } = require("uuid");

function generateOrder(products) {
  const end = Math.floor(Math.random() * products.length);

  const rnd = Math.floor(Math.random() * products.length);

  const start = rnd < end ? rnd : 0;

  const subset = products.slice(start, end);

  const orderProducts = subset.map((product) => ({
    ...product,
    quantity: Math.floor(Math.random() * 100),
  }));

  const order = {
    id: uuidV4(),
    deliveryAddress: {
      street: "Rottenbucher-Str",
      number: 24,
      zip: 827382,
    },
    customerInfo: {
      name: "John Doe",
    },
    products: orderProducts,
  };

  return order;
}

module.exports = {
  generateOrder,
};
