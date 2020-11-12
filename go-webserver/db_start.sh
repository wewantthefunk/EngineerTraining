#!/bin/bash
docker stop some-postgres
docker stop go-webserver_web_1
docker rm some-postgres
docker run --name some-postgres --network my_network -p 5432:5432 -v $(pwd):/test -e POSTGRES_PASSWORD=mysecretpassword -d postgres

echo 10
sleep 1
echo 9
sleep 1
echo 8
sleep 1
echo 7
sleep 1
echo 6
sleep 1
echo 5
sleep 1
echo 4
sleep 1
echo 3
sleep 1
echo 2
sleep 1
echo 1
sleep 1
docker exec some-postgres chmod 777 /test/execute.sh
docker exec some-postgres /test/execute.sh
docker-compose up --build
