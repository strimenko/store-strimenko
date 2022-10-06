run:
	sudo docker compose up --build -d

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up

stop:
	docker compose down