FROM golang:alpine

ENV GOPATH /app
ENV PATH $PATH:$GOPATH/bin

RUN echo $PATH

RUN apk update
RUN apk add git

#golang dep replaces govendor
RUN go get -u github.com/golang/dep/cmd/dep

#RUN ls /app/src/fake_news
COPY . /app/src/fake_news/

WORKDIR /app/src/fake_news

RUN dep ensure

WORKDIR app

RUN go install -v fake_news/sentiment
RUN ls $GOPATH/bin
CMD $GOPATH/bin/sentiment

EXPOSE 8000