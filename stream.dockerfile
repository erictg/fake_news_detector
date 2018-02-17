FROM alpine

COPY bin/stream_api main


CMD main

EXPOSE 8001
