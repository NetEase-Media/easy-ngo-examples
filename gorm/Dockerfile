FROM mysql

ENV MYSQL_ROOT_PASSWORD=root

WORKDIR /app

COPY ./setup.sql /app/setup.sql

RUN mv /app/setup.sql /docker-entrypoint-initdb.d/setup.sql

EXPOSE 3306