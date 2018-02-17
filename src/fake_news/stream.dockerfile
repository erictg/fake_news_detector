FROM golang:alpine

ENV GOPATH /app
ENV PATH $PATH:$GOPATH/bin

RUN echo $PATH

RUN apk update
RUN apk add git

#golang dep replaces govendor
RUN go get -u github.com/golang/dep/cmd/dep

#RUN ls /app/src/fact_check
COPY . /app/src/fact_check/

WORKDIR /app/src/fact_check

RUN dep ensure

WORKDIR app

RUN go install -v fake_news/stream_api
CMD bin/stream_api

EXPOSE 8001
