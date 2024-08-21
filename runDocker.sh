#!/bin/bash

NAME="testPostgres"
PORT="5432"
PASSWORD="0000"
USER="postgres"

docker run -d --name "$NAME" -p "$PORT":"$PORT" -e POSTGRES_PASSWORD="$PASSWORD" -e POSTGRES_USER="$USER" mypostgresql:latest