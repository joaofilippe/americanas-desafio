FROM golang:latest as builder
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/ -o ../server

ENTRYPOINT [ "/app/server" ]

FROM scratch
COPY --from=builder /app/server /server
ENTRYPOINT [ "/server" ]