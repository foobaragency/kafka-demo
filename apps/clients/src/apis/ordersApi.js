const { ordersClient } = require("./clients");

async function createOrder(order) {
  try {
    await ordersClient.post("/", order);
  } catch (error) {}
}

module.exports = {
  createOrder,
};
