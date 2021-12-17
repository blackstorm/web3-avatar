FROM golang:1.17.5-buster as BUILDER

WORKDIR /app
ADD ./* /app/
RUN make build

FROM debian:buster

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates && apt-get clean
COPY --from=BUILDER /app/main .

EXPOSE 8080 8080

ENTRYPOINT ["/app/main"]