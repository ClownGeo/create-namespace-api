FROM golang:1.26-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ns-manager ./cmd/server/main.go


FROM scratch

WORKDIR /
COPY --from=builder /app/ns-manager /opt/ns-manager

EXPOSE 8080

CMD ["/opt/ns-manager"]