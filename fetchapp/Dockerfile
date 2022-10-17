# build stage
FROM golang:1.17.13-alpine as builder
LABEL maintainer="Mohamad Taufiq <mohamad.taufiq41186@gmail.com>"
ARG FETCHAPP_SVC_NAME
ARG FETCHAPP_SVC_PORT
RUN apk --no-cache add git curl build-base
COPY fetchapp/. /go/src/$FETCHAPP_SVC_NAME
RUN ls -lah /go/src/$FETCHAPP_SVC_NAME
WORKDIR /go/src/$FETCHAPP_SVC_NAME

# Prefer using go module
# make sure go.mod and go.sum exists by running go mod init
# Uncomment line below if you've configured go module
# RUN go get -v

# Prefer using dep
# make sure Gopkg.lock and Gopkg.toml exists by running dep ensure -v
# Uncomment line below if you've configured go module
# RUN dep ensure -v --vendor-only

# Because by default is using dep, then this command will try 2 options, go mod then dep
# Comment line below if you already decide your vendoring management
RUN ((go mod download && go mod tidy) || dep ensure -v --vendor-only)

# Run unit test
# RUN go test -v ./...

# Please make sure where your package main is, and specifically, where you main function is
# it could be on cmd/main.go or main.go or somewhere else
RUN go build -o /go/bin/$FETCHAPP_SVC_NAME ./src/cmd

# final stage
FROM alpine:latest
LABEL maintainer="Mohamad Taufiq <mohamad.taufiq41186@gmail.com>"
ARG FETCHAPP_SVC_PORT
ARG FETCHAPP_SVC_NAME
ENV GOPATH="/go"
RUN apk --no-cache add tzdata && echo "Asia/Jakarta" > /etc/timezone
RUN printf "#!/bin/sh\n\nwhile true; do\n\techo \"[INFO] Starting Service at \$(date)\"\n\t(./$FETCHAPP_SVC_NAME >> ./history.log || echo \"[ERROR] Restarting Service at \$(date)\")\ndone" > run.sh
RUN printf "#!/bin/sh\n./run.sh & tail -F ./history.log" > up.sh
RUN chmod +x up.sh run.sh
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/bin/$FETCHAPP_SVC_NAME .
COPY --from=builder /go/src/$FETCHAPP_SVC_NAME/etc/conf/conf.yaml /etc/conf/conf.yaml
# COPY --from=builder /go/src/$FETCHAPP_SVC_NAME/demo.db .
EXPOSE ${FETCHAPP_SVC_PORT}
CMD ["./up.sh"]