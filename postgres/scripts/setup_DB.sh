# docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=123 postgres:12-alpine
docker run  -p 5432:5432 \
    -e POSTGRES_PASSWORD=123 \
    -e POSTGRES_USER=sandbox_admin \
    -e POSTGRES_DB=sandbox_db1 \
    -e PGDATA=/var/lib/postgresql/data/pgdata \
    -v $PWD/../pgdata/:/var/lib/postgresql/data \
    postgres:12-alpine

# sleep 5

# pgcli postgres://sandbox_admin:123@127.0.0.1:5432/sandbox_admin_db1
#psql postgres://sandbox_admin:123@127.0.0.1:5432/sandbox_admin_db1 < init_db.sql

