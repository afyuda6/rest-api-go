FROM golang:1.23.4

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends \
    wget \
    curl \
    ca-certificates && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /rest-api-go

COPY . /rest-api-go

RUN go mod tidy && go build -o rest-api-go .

EXPOSE 8080

CMD ["./rest-api-go"]