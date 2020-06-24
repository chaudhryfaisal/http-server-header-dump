############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
WORKDIR $GOPATH/src/app
COPY main.go .
# Using go get.
RUN go get -d -v
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app
############################
# STEP 2 build a small image
############################
FROM alpine
EXPOSE 8080
ENV PORT=8080
ENV DUMP_ENV=false
ENV ENV_MATCH=".*"
COPY --from=builder /go/bin/app /app
ENTRYPOINT /app --port=$PORT --includeEnv=$DUMP_ENV --matchEnv=$ENV_MATCH