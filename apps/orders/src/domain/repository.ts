import { createOrderEvents } from "~infra/events"
import { Order } from "./model"
import { v4 as uuidV4 } from "uuid"
import { mongoDB } from "~infra/mongodb"

export function createOrder(order: Order): Order {
  const newOrder = {
    id: uuidV4(),
    ...order,
  }

  createOrderEvents(newOrder)

  return newOrder
}

export async function checkProductAvailability(order: Order): Promise<boolean> {
  const ids = order.products.map((p) => p.id)

  const client = await mongoDB()

  const col = client.db().collection("product_stock")

  const results = await col.find({
    _id: {
      $in: ids,
    },
  })

  let available = true

  await results.forEach((p) => {
    const product = order.products.find((op) => op.id === p._id.toString())
    const availableStock = p["CURRENT_STOCK"]

    if (availableStock && product.quantity > availableStock) {
      available = false
    }
  })

  return available
}
