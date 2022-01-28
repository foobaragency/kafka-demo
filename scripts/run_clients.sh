#!/bin/bash
trap "exit" INT TERM ERR
trap "kill 0" EXIT

cd apps/clients && yarn dev &

wait