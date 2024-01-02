psql -U worker -c "create database recordings"
psql -U worker -d recordings -f /init-table/create-tables.sql
psql -U worker -d recordings -c "SELECT * FROM album"
