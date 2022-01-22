import { Order, OrderProduct } from "~domain/model"
import { v4 as uuidV4 } from "uuid"
import { CloudEvent, CONSTANTS, HTTP } from "cloudevents"

function mapProducts(order: Order): CloudEvent<OrderProduct>[] {
  return order.products.map((product) => {
    const attributes = {
      "ce-specversion": "1.0",
      "ce-type": "kafka_demo.BindProductToOrder",
      "ce-source": "ordersAPI/v1/orders",
      "ce-id": uuidV4(),
      "ce-time": new Date().toISOString(),
      "Content-Type": CONSTANTS.MIME_JSON,
      "ce-destination": "order_products",
    }

    const productData = {
      key: order.id,
      ...product,
    }

    const data = {
      ...attributes,
      ...productData,
    }

    const event = HTTP.toEvent({ headers: attributes, body: data })

    return new CloudEvent({
      ...event,
      data,
    })
  })
}

export function mapOrderToEvents(
  order: Order
): CloudEvent<Order | OrderProduct>[] {
  const attributes = {
    "ce-specversion": "1.0",
    "ce-type": "kafka_demo.CreateOrder",
    "ce-source": "ordersAPI/v1/orders",
    "ce-id": uuidV4(),
    "ce-time": new Date().toISOString(),
    "Content-Type": CONSTANTS.MIME_JSON,
    "ce-destination": "orders",
  }

  const orderData = {
    key: order.id,
    ...order,
  }

  const data = {
    ...attributes,
    ...orderData,
  }

  const event = HTTP.toEvent({ headers: attributes, body: data })

  const orderEvent = new CloudEvent({
    ...event,
    data,
  })

  const productEvents = mapProducts(order)

  return [orderEvent, ...productEvents]
}
