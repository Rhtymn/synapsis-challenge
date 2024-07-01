FROM golang:1.22.4
WORKDIR /application
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /synapsis-challenge

ARG SERVER_ADDR
ENV SERVER_ADDR=${SERVER_ADDR}

EXPOSE 8080
CMD ["/synapsis-challenge"]