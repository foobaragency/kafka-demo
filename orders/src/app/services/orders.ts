import { Request, Response } from "express"
import { createOrder } from "~domain/repository"

export function createOrderHandler(req: Request, res: Response) {
  createOrder(req.body)

  res.status(200).end()
}
