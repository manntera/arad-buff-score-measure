FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src
COPY . .
RUN go build ./cmd/server/main.go

# runtime image
FROM alpine
COPY --from=builder /go/src /app
COPY --from=builder /go/src/setting /app/setting

CMD /app/main $PORT