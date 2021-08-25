FROM golang:1.16 as builder
WORKDIR /app
COPY hello.go ./
RUN go build -o hello hello.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/hello ./
CMD ["./hello"]
