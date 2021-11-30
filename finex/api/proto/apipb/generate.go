package apipb

//go:generate protoc -I=.. -I=../../../ent/proto --go_out=.. --go-grpc_out=.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative apipb/asset.proto
