FROM clearlinux/golang:1.15 as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server ./cmd/server

FROM clearlinux:latest

WORKDIR /dist

COPY --from=builder /build/server .

CMD ["./server"]
