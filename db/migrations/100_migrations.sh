#!/bin/bash
# golang-migrate を使用して初回起動時にマイグレーションを実行するスクリプト

MYSQL_USER="docker"
MYSQL_PASSWORD="docker"

OP_DATABASES=()

for OP_DATABASE in "${OP_DATABASES[@]}" ; do
    echo "Migrate ${OP_DATABASE}..."

    MYSQL_DSN="mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@unix(/var/run/mysqld/mysqld.sock)/${OP_DATABASE}"
    migrate -path /migrations -database "${MYSQL_DSN}" up
done
