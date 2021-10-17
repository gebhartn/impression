#!/bin/sh
chmod o+w ./database
docker-compose up --build
docker-compose down
rm -f ./database/impress.db
