FROM golang:1.18-alpine as builder
RUN apk add git
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/src
RUN go build -o personalFinanceManager

FROM alpine

WORKDIR /root/

COPY --from=builder /app/src/personalFinanceManager .
COPY --from=builder /app/src/*.properties .

CMD ["./personalFinanceManager"]