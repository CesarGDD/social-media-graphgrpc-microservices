go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc commentspb/commentspb.proto --go_out=. --go-grpc_out=.