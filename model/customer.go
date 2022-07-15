package model

type Customer struct {
	LopeiId int32
	Balance float32
}

/*
protoc --go_out=./service --go-grpc_out=./service model/lopei.proto
*/
