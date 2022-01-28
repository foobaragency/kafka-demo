import { MongoClient } from "mongodb"

const URI = process.env.MONGODB_URI

export const mongoClient = new MongoClient(URI)

export async function mongoDB() {
  const client = await mongoClient.connect()
  return client
}
