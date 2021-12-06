.PHONY: docker/build
docker/build:
	@echo 'Building facilittei/ecomm image'
	docker build --no-cache -t facilittei/ecomm .

.PHONY: docker/run
docker/run:
	@echo 'Starting container on port :4000'
	docker run -d -p 4000:80 --name ecomm facilittei/ecomm

.PHONY: docker/remove
docker/remove:
	@echo 'Stopping container'
	docker stop ecomm
	@echo 'Removing container'
	docker rm ecomm
	@echo 'Removing image'
	docker rmi facilittei/ecomm

.PHONY: compose/up
compose/up:
	@echo 'Starting services via Docker compose'
	docker-compose up -d

.PHONY: compose/down
compose/down:
	@echo 'Stopping services via Docker compose'
	docker-compose down

.PHONY: compose/buildup
compose/buildup:
	@echo 'Re-building image and starting services via Docker compose'
	docker-compose up --build -d

.PHONY: test/all
test/all:
	@echo 'Starting application tests'
	go test -v ./...
	@echo 'Tests have been finished'