export tag=v1.0
root:
	export ROOT=https://github.com/psjbasil/httpserver

build:
	mkdir -p bin/amd64
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/amd64/app

release: build
	echo "building httpserver image"
	docker build -t psjbasil/httpserver:${tag} .

push: release
	echo "pushing to psjbasil/httpserver"
	docker push psjbasil/httpserver:${tag}
pull:
	sudo docker pull psjbasil/httpserver:${tag}
run:
	sudo docker run psjbasil/httpserver:${tag}