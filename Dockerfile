FROM golang:alpine AS builder
RUN apk update
WORKDIR /

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o a.out

FROM alpine as runner
RUN apk update
WORKDIR /

COPY --from=builder /a.out .

CMD ["./a.out"]
