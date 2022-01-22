import { Order } from "~domain/model"
import { mapOrderToEvents } from "~util/mapOrderToEvents"

export function createOrderEvents(order: Order) {
  const events = mapOrderToEvents(order)

  events.forEach((event) => event.emit())
}
