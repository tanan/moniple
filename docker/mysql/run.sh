#!/bin/sh

tag=5.7.21

docker stop moniple-db
docker rm -v moniple-db
docker run -i -t -d -p 3308:3306 -e MYSQL_ROOT_PASSWORD=xxxx --name moniple-db moniple-db:${tag}
 
