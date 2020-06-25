NETWORK := simplegrpc
CONTAINER := simplegrpc
CLIENT := simplegrpc-client
SERVER := simplegrpc-server

protobuf:
	protoc -I messenger/ messenger/messenger.proto --go_out=plugins=grpc:messenger --go_opt=paths=source_relative

build:
	docker build -t $(CONTAINER) .

run-server:
	docker run -it -v `pwd`:/home/dev --net $(NETWORK) --name $(SERVER) $(CONTAINER) 
run-client:
	docker run -it -v `pwd`:/home/dev --net $(NETWORK) --name $(CLIENT) $(CONTAINER)

network:
	docker network create $(NETWORK)

clean:
	docker rm -f $(SERVER) || true
	docker rm -f $(CLIENT) || true 
	docker network rm $(NETWORK) || true
