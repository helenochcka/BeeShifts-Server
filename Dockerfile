FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /beeshifts-server ./cmd/gin

ENV CONFIG_PATH="config/test_config.yaml"

CMD ["/beeshifts-server"]