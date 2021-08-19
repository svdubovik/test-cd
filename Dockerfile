FROM golang:1.16 as builder
COPY hello.go ./
RUN go build -o hello hello.go

FROM scratch
COPY --from=builder /go/hello ./
CMD ["./hello"]
