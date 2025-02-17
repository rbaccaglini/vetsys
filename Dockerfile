FROM amd64/golang:1.23 AS BUILDER

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY config config
COPY docs docs
COPY internal internal
COPY pkg pkg
COPY test test

WORKDIR /app/cmd/simple_crud_golang

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/vetsys .

FROM golang:1.23-alpine AS RUNNER

WORKDIR /root/

COPY --from=BUILDER /app/vetsys .

EXPOSE 8080

CMD ["./vetsys"]
