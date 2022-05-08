FROM golang:1.18-alpine as builder
RUN apk add git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build

FROM alpine

WORKDIR /root/

COPY --from=builder /app/personalFinanceManager .

CMD ["./personalFinanceManager"]