FROM node:alpine

WORKDIR /root
COPY frontend/ /root/frontend
COPY Makefile /root
RUN apk add --update make
RUN make prod

FROM golang:alpine

WORKDIR /usr/local/go/src/github.com/rdesousa/dashboard-gitlab
RUN apk add --update make git
RUN go get github.com/gin-gonic/gin
COPY Makefile .
COPY ./main.go .
COPY api/ api
RUN make build
RUN chmod +x dashboard-gitlab

FROM alpine:latest
WORKDIR /root/
COPY test.json /tmp
RUN apk add --update libcurl
COPY --from=0 /root/frontend frontend
COPY --from=1 /usr/local/go/src/github.com/rdesousa/dashboard-gitlab/dashboard-gitlab .
CMD ./dashboard-gitlab
