# build stage
FROM node:14.17.0-alpine3.13 as builder
LABEL maintainer="Mohamad Taufiq <mohamad.taufiq41186@gmail.com>"
ARG AUTHAPP_SVC_NAME
ARG AUTHAPP_SVC_PORT
RUN apk --no-cache add git curl build-base sqlite
COPY authapp/. /$AUTHAPP_SVC_NAME
WORKDIR /$AUTHAPP_SVC_NAME

RUN npm install
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
# RUN ((go mod download && go mod tidy) || dep ensure -v --vendor-only)

# Run unit test
# RUN go test -v ./...

# RUN npm run start
# Please make sure where your package main is, and specifically, where you main function is
# it could be on cmd/main.go or main.go or somewhere else
# RUN go build -o /go/bin/$SVC_NAME ./src/cmd

# final stage
FROM node:14.17.0-alpine3.13
LABEL maintainer="Mohamad Taufiq <mohamad.taufiq41186@gmail.com>"
ARG AUTHAPP_SVC_PORT
ARG AUTHAPP_SVC_NAME
RUN apk --no-cache add tzdata && echo "Asia/Jakarta" > /etc/timezone
RUN printf "#!/bin/sh\n\nwhile true; do\n\techo \"[INFO] Starting Service at \$(date)\"\n\t(npm run start >> ./history.log || echo \"[ERROR] Restarting Service at \$(date)\")\ndone" > run.sh
RUN printf "#!/bin/sh\n./run.sh & tail -F ./history.log" > up.sh
RUN chmod +x up.sh run.sh
COPY --from=builder /$AUTHAPP_SVC_NAME .
EXPOSE ${AUTHAPP_SVC_PORT}
CMD ["./up.sh"]