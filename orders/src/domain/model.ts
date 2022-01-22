export type Product = {
  id: string
  name: string
  stock: number
}

export type OrderProduct = {
  id: Product["id"]
  quantity: number
}

export type DeliveryAddress = {
  street: string
  number: number
  zip: number
}

export type CustomerInfo = {
  name: string
}

export type Order = {
  id: string
  products: OrderProduct[]
  deliveryAddress: DeliveryAddress
  customerInfo: CustomerInfo
}
