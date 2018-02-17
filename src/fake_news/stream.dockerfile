FROM golang:alpine

ENV GOPATH /app
ENV PATH $PATH:$GOPATH/bin
ENV GOBIN $GOPATH/bin
ENV DOCKER_INSIDE yee
RUN echo $PATH

RUN apk update
RUN apk add git

#golang dep replaces govendor
RUN go get -u github.com/golang/dep/cmd/dep

COPY ./stream_api /app/src/fact_check
COPY Gopkg.lock /app/src/fact_check/Gopkg.lock
COPY Gopkg.toml /app/src/fact_check/Gopkg.toml
WORKDIR /app/src/fact_check

RUN dep ensure

RUN go install -v main.go

CMD $GOBIN/main

EXPOSE 8001
