import { Headers } from "cloudevents"

export function replaceCEDash(attributes: Headers) {
  const entries = Object.entries(attributes)
    .filter(([key]) => !["Content-Type", "ce-destination"].includes(key))
    .map(([key, value]) => {
      return [key.replace("-", "_"), value]
    })

  return Object.fromEntries(entries)
}
