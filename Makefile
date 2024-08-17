
docker/build:
	docker build -t fam-chat .

docker/run:
	docker run -p 8080:8080 fam-chat
