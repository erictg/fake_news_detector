FROM openresty/openresty:latest

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf

COPY default.conf /usr/local/openresty/nginx/conf/conf.d/default.conf
COPY ./web /usr/local/openresty/nginx/html/

RUN apk update && apk add openssl
COPY zuul.key /etc/ssl/zuul.key
COPY zuul.csr /etc/ssl/zuul.csr
COPY zuul.crt /etc/ssl/zuul.crt

ENTRYPOINT ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]