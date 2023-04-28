FROM mysql:5.7

ENV MYSQL_DATABASE=testdb \
    MYSQL_USER=mysql_user \
    MYSQL_PASSWORD=root \
    MYSQL_ROOT_PASSWORD=root

COPY ./init.sql /docker-entrypoint-initdb.d/