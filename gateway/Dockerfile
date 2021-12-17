FROM golang:1.17.5-buster as BUILDER

WORKDIR /app
ADD ./* /app
RUN make build

FROM debian:buster

WORKDIR /app
COPy --from=BUILDER /app/main .

ENTRYPOINT ["/app/main"]