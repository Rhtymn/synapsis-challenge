FROM golang:1.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o synapsis-challenge-be .

FROM alpine:latest
COPY --from=builder /app/synapsis-challenge-be /synapsis-challenge-be
EXPOSE 8080
CMD ["/synapsis-challenge-be"]
