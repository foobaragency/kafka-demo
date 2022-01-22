import { Request, Response, NextFunction } from "express"
import { Emitter, emitterFor, Mode } from "cloudevents"
import { run } from "./kafka"

const emit = emitterFor(run, { mode: Mode.BINARY })

export function emitter(req: Request, res: Response, next: NextFunction) {
  Emitter.on("cloudevent", emit)
  next()
}
