#!/bin/bash -eu

while [ true ]
do
    echo "${DB_USERNAME}" > /data/DB_USERNAME
    echo "${DB_PASSWORD}" > /data/DB_PASSWORD
    sleep 10
done
