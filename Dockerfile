FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod ./
COPY *.go ./
COPY piano_interval_fetch.html ./
COPY static ./static

RUN go build -o /server server.go findinterval.go

FROM alpine:latest

WORKDIR /root/

COPY --from=build /server ./
COPY --from=build /app/piano_interval_fetch.html ./
COPY --from=build /app/static ./static

EXPOSE 3000

CMD ["./server"]