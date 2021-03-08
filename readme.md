docker build -t go-heartbeat-service .

docker run -p 8080:8080 go-heartbeat-service

sudo docker exec -it go-heartbeat-service pwd

sudo docker exec -it --entrypoint /bin/sh go-heartbeat-service