const axios = require("axios");

const PRODUCTS_URI = process.env.PRODUCTS_URI;
const ORDERS_URI = process.env.ORDERS_URI;

const productsClient = axios.create({
  baseURL: `${PRODUCTS_URI}/v1/products`,
  timeout: 2000,
});

const ordersClient = axios.create({
  baseURL: `${ORDERS_URI}/v1/orders`,
  timeout: 2000,
});

module.exports = {
  productsClient,
  ordersClient,
};
