# simple-grpc
Simple gRPC app

## Running
1. Build the container with cached dependencies
`make build`
2. Make the network 
`make network`
3. Run the server
`make run-server`
4. Run the client
`make run-client`
5. Clean up resources
`make clean`

## Additional
To update messenger/messenger.pb.go run `make protobuf`