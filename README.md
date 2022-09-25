# clean-architecture

// protobuf 自動生成コマンド
protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/<ファイル名>.proto
