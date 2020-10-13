FROM golang:latest AS build

WORKDIR /work_dir

# first download dependencies
COPY go.mod /work_dir
COPY go.sum /work_dir
RUN go mod download

# then copy source code
COPY / /work_dir


RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /storesservice ./cmd/storesservice


FROM alpine:latest

WORKDIR /

COPY --from=build /storesservice /storesservice/

WORKDIR /storesservice

RUN chmod +x ./storesservice

EXPOSE 8080

CMD ["./storesservice"]