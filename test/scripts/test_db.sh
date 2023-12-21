#!/bin/bash

#docker run --name postgresdb2 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Password.1 -p 6432:5432 -d postgres

#echo "Postgresdb starting..."
#sleep 3

docker exec -it postgresdb psql -U postgres -d postgres -c "CREATE DATABASE go_productapp"
sleep 3
echo "Database go_productapp created"

docker exec -it postgresdb psql -U postgres -d go_productapp -c "
create table if not exists products
(
    id bigserial not null primary key,
    name varchar(255) not null,
    price double precision not null,
    discount double precision,
    store varchar(255) not null
);
"

sleep 3
echo "Table products created"

