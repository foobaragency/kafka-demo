import express from "express"
import bodyParser from "body-parser"
import router from "~services/router"
import { emitter } from "~infra/emitter"

const app = express()
const PORT = 8081

app.use(emitter)

app.use(bodyParser.json())

app.use("/v1", router)

app.listen(PORT, () => {
  console.log(`server started at http://localhost:${PORT}`)
})
