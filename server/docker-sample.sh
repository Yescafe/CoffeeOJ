#!/usr/bin/env bash

docker run --name some-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -d mysql
docker run --name some-redis -p 6379:6379 -d redis

