FROM mysql:5.7.21

ADD init.sh /docker-entrypoint-initdb.d/.

RUN mkdir -p /data/init.d
ADD init.d /data/init.d/.
ADD my.cnf /etc/mysql/.

EXPOSE 3306