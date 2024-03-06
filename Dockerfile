FROM golang:1.21.1-alpine3.18 AS BuildStage
WORKDIR /
COPY . .
RUN CGO_ENABLED=0 go build -o helloworld

FROM alpine:latest
WORKDIR /app
COPY --from=BuildStage helloworld /app/helloworld
ENTRYPOINT ["./helloworld"]
