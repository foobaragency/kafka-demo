import { Message } from "cloudevents"
import { Kafka, IHeaders } from "kafkajs"

const KAFKA_USER = process.env.KAFKA_USER
const KAFKA_PASSWORD = process.env.KAFKA_PASSWORD
const KAFKA_BROKERS = process.env.KAFKA_BROKERS

const kafka = new Kafka({
  clientId: "ordersAPI",
  brokers: [KAFKA_BROKERS],
  connectionTimeout: 5_000,
  ssl: true,
  sasl: {
    username: KAFKA_USER,
    password: KAFKA_PASSWORD,
    mechanism: "plain",
  },
})

const producer = kafka.producer()

export async function run(message: Message) {
  await producer.connect()

  const body = JSON.parse(message.body as string)

  const key = body.key as string
  const value = message.body as Buffer
  const headers = message.headers as IHeaders

  await producer.send({
    topic: message.headers["ce-destination"] as string,
    messages: [{ key, value, headers }],
  })
}
