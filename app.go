package main

import "lopei-grpc-server/delivery"

func main() {
	delivery.Server().Run()
}

// set GRPC_URL=localhost:8888
