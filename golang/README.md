```bash
# create shared network
docker network create demonet

# start pg
docker run --net demonet \
 --rm \
 -d \
 --name demo-postgres \
 -e POSTGRES_PASSWORD=mysecretpassword \
 postgres:9.6

sleep 2
DB_URL="postgres://postgres:mysecretpassword@demo-postgres:5432/postgres?sslmode=disable";

# create application table
docker run --net demonet \
 --rm \
 -i \
 postgres:9.6 \
 psql "${DB_URL}" -f - \
 << EOSQL
/* Create Application Table */
CREATE TABLE pets (
  id serial primary key,
  name varchar(256)
);
EOSQL

# build application
docker build -t pet-store:latest .

# start application
docker run --net demonet \
 --rm \
 --name demo-app \
 -e DB_URL="${DB_URL}" \
 -p 8080:8080 \
 pet-store:latest

# access the application
curl \
 -i \
 -d @- \
 -H "Content-Type: application/json" \
 localhost:8080/pet \
 << EOL
{
  "name": "Cuteness Overload!"
}
EOL

curl -i localhost:8080/pets

# clean up
docker rm -f demo-postgres demo-app
docker network rm demonet
```