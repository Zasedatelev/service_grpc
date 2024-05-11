FROM golang:1.22.0-alpine3.19 as  builder 

COPY . /github.com/Zasedatelev/service_go/grpc/sourse/
WORKDIR /github.com/Zasedatelev/service_go/grpc/sourse/


RUN go mod download
RUN go build -o ./bin/test_grpc_server cmd/grpc_server/main.go


FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Zasedatelev/service_go/grpc/sourse/bin/test_grpc_server .

CMD ["./test_grpc_server"]