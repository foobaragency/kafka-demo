import { createOrderEvents } from "~infra/events"
import { Order } from "./model"
import { v4 as uuidV4 } from "uuid"

export function createOrder(order: Order) {
  createOrderEvents({
    id: uuidV4(),
    ...order,
  })
}
