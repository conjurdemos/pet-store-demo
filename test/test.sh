#!/bin/bash

cd $(dirname $0);

if [[ $# -ne 1 ]]
then
        echo "Usage : $0 DATABASE_TYPE"
        exit
fi

case "$1" in
postgres) export DB_TYPE=postgres
          export DB_URL=postgresql://postgres:5432/postgres
          export DB_USERNAME=postgres
          ;;
mysql)  export DB_TYPE=mysql
        export DB_URL=mysql://mysql:3306/mysql
        export DB_USERNAME=mysql
        ;;
*) echo "DATABASE_TYPE $1 is not supported"
   exit 1
   ;;
esac

cleanup () {
  docker-compose down -v;
}
trap cleanup EXIT QUIT INT;

cleanup;
docker-compose up -d ${DB_TYPE};

echo "Waiting for $DB_TYPE to start"
while ! docker-compose ps ${DB_TYPE} | grep healthy > /dev/null 2>&1;
  do
    >&2 printf '. '
    sleep 1
done

echo ""
>&2 echo "$DB_TYPE is up - continuing"

docker-compose up -d app test;

echo "Waiting for app"
while ! docker-compose exec -T test curl -v app:8080 > /dev/null 2>&1;
  do
    >&2 printf '. '
    sleep 1
done

echo ""
>&2 echo "app is up - continuing"

echo ""
echo 'Creating pet:'
echo '{"name": "Mr. Init"}'
docker-compose exec -T test curl -s -d '{"name": "Mr. Init"}' -H "Content-Type: application/json" app:8080/pet

echo "Retrieving pet:"
docker-compose exec -T test curl -s app:8080/pets
echo ""
echo ""
