FROM golang:1.25.1-alpine3.22 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/app/main.go

FROM alpine:3.22.3
WORKDIR /app
COPY --from=builder /app/main .

ENV SERVER_ADDR="localhost:8082" 
ENV PSQL_HOST="localhost"
ENV PSQL_PORT="5432"
ENV PSQL_USER="postgres"
ENV PSQL_PASSWORD="postgres"
ENV PSQL_DBNAME="pdd-datastore"
ENV PSQL_SSLMODE="disable"
ENV PSQL_CONNECT_TIMEOUT=10
ENV PSQL_CONNECT_WAIT_TIME=3
ENV PSQL_CONNECT_ATTEMPTS=3
ENV PSQL_CONNECT_BLOCKS=false
ENV PSQL_CLOSE_TIMEOUT=10

ENV GITHUB_API_TOKEN=""
ENV GITHUB_API_ENDPOINT="https://api.github.com/graphql"

EXPOSE 8080
CMD [ "/app/main"]