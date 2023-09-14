
postgres:
	docker run --name filtronicdb -p 5432:5432 -e POSTGRES_USER=filtronic -e POSTGRES_PASSWORD=secret -d postgres:16beta3

createdb:
	docker exec -it filtronicdb createdb --username=filtronic --owner=filtronic edms