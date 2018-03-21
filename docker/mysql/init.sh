#!/bin/sh

mysql="mysql --default-character-set=utf8"
mysql_root="${mysql} -uroot -pxxxx"

${mysql_root} < /data/init.d/schema.sql
