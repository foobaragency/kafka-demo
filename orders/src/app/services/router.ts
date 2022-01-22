import { Router } from "express"
import { createOrderHandler } from "./orders"

const router = Router()

router.route("/orders").post(createOrderHandler)

export default router
