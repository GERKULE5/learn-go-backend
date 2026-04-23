up:
	docker-compose up -d crud-database
	go run .\cmd\api\main.go
down: 
	docker-compose down crud-database