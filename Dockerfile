FROM golang:latest as builder
WORKDIR /app
COPY . .

ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_USER=postgres
ENV DB_PASSWORD=12345678
ENV DB_NAME=list_node
ENV DB_SSLMODE=disable

RUN CGO_ENABLED=0 GOOS=linux go build -C cmd/ -o ../server
ENTRYPOINT [ "/app/server" ]


# FROM scratch
# COPY --from=builder /app/server /server