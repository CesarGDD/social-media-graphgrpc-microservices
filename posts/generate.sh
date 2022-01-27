go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
protoc postspb/postspb.proto --go_out=. --go-grpc_out=.