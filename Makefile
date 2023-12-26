start:
	docker-compose -f ./build/docker-compose.yml up --abort-on-container-exit --force-recreate --build
