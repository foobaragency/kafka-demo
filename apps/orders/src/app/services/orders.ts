import { Request, Response } from "express"
import { checkProductAvailability, createOrder } from "~domain/repository"

export async function createOrderHandler(req: Request, res: Response) {
  const available = await checkProductAvailability(req.body)

  if (!available) {
    return res
      .status(500)
      .json({ message: "One or more products out of stock" })
  }

  const order = createOrder(req.body)

  return res.status(200).json(order)
}
