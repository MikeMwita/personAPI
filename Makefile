
postgres:
	docker run --name filtronicdb -p 5432:5432 -e POSTGRES_USER=filtronic -e POSTGRES_PASSWORD=secret -d postgres:16beta3

createdb:
	docker exec -it filtronicdb createdb --username=filtronic --owner=filtronic edms


network:
	docker network create personapi
attach:
	docker run --name filtronicdb -p 5433:5433 -e POSTGRES_USER=filtronic -e POSTGRES_PASSWORD=secret -d --network personapi postgres:16beta3





