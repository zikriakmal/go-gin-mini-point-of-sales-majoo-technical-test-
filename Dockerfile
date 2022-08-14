FROM golang:alpine as builder
WORKDIR /go/src/majoo
COPY . .
RUN go mod download
RUN go build main.go
EXPOSE 9090

#FROM alpine
#WORKDIR /app
#COPY --from=builder /app/tts /gcapiv2

