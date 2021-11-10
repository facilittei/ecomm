.PHONY: docker/build
docker/build:
	@echo 'Building facilittei/ecomm image'
	docker build --no-cache -t facilittei/ecomm .

.PHONY: docker/run
docker/run:
	@echo 'Starting container on port :4000'
	docker run -d -p 4000:4000 --name ecomm facilittei/ecomm

.PHONY: docker/cleanup
docker/cleanup:
	@echo 'Stopping container'
	docker stop ecomm
	@echo 'Removing container'
	docker rm ecomm
	@echo 'Removing image'
	docker rmi facilittei/ecomm
