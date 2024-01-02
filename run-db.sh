docker volume create 03-go-db
docker compose up -d
docker exec -i pg-database psql -U worker -c "create database recordings"
docker exec -i pg-database psql -U worker -d recordings -f /init-table/create-tables.sql
docker exec -i pg-database psql -U worker -d recordings -c "SELECT * FROM album"
