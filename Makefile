#!/bin/bash
build-dev:
	docker build . -f DockerfileDev -t motecshine:dev

run-dev:
	docker run -p 8801:8801 motecshine:dev
