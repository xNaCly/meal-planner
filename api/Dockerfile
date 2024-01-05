FROM alpine:latest as builder
WORKDIR /api
COPY . .
RUN apk add --no-cache go
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o server

FROM scratch
COPY --from=builder /api/server ./server
CMD ["./server"]
