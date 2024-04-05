DOCKER_IMAGE = "go-algorithm"

docker_run:
	docker run -it --rm --mount type=bind,source=.,target=/app -w /app ${DOCKER_IMAGE}
