#!/bin/bash

#docker run --name postgresdb2 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=Password.1 -p 6432:5432 -d postgres

#echo "Postgresdb starting..."
#sleep 3

#docker exec -it postgresdb2 psql -U postgres -d postgres -c "CREATE DATABASE muratcoskun_productapp"
#sleep 3
#echo "Database muratcoskun_productapp created"

#docker exec -it postgresdb2 psql -U postgres -d muratcoskun_productapp -c "
# create table if not exists products
# (
#     id bigserial not null primary key,
#     name varchar(255) not null,
#     price double precision not null,
#     discount double precision,
#     store varchar(255) not null
# );
# "

# sleep 3
# echo "Table products created"

