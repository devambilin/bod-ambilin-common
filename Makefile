generate-proto:
	protoc ${file} -I.  --micro_out=${GOPATH}/src/ --go_out=${GOPATH}/src/
server:
	go run cmd/main.go