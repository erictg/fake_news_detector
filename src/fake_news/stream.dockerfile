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

RUN ls /app/src/fact_check
COPY ./stream_api /app/src/fact_check/stream_api

COPY Gopkg.lock /app/src/fact_check/Gopkg.lock
COPY Gopkg.toml /app/src/fact_check/Gopkg.toml
WORKDIR /app/src/fact_check/stream_api

RUN dep ensure

RUN go install -v fake_news/stream_api
CMD $GOBIN/main

EXPOSE 8001
