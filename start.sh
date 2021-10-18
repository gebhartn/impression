#!/bin/sh
chmod o+w ./database
docker-compose --env-file ./env up --build
docker-compose down
rm -f ./database/impress.db
