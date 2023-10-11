docker-start:
	docker-compose up --build --remove-orphans -d

docker-stop:
	docker-compose down --volumes --remove-orphans
	docker container prune --force