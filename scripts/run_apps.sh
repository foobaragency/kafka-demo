#!/bin/bash
trap "exit" INT TERM ERR
trap "kill 0" EXIT

cd apps/products && gow run internal/products/app/server.go &
cd apps/orders && yarn dev &

wait