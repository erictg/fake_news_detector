FROM openresty/openresty:latest

COPY nginx.conf /usr/local/openresty/nginx/conf/nginx.conf

COPY default.conf /usr/local/openresty/nginx/conf/conf.d/default.conf
COPY ./web /usr/local/openresty/nginx/html/

RUN apk update && apk add openssl
RUN /usr/bin/openssl genrsa -out /etc/ssl/zuul.key 1024
RUN /usr/bin/openssl req  -new -newkey rsa:4096 -days 365 -nodes -subj "/C=/ST=/L=/O=/CN=zuul" -keyout /etc/ssl/zuul.key -out /etc/ssl/zuul.csr
RUN /usr/bin/openssl x509 -req -days 365 -in /etc/ssl/zuul.csr -signkey /etc/ssl/zuul.key -out /etc/ssl/zuul.crt

ENTRYPOINT ["/usr/local/openresty/bin/openresty", "-g", "daemon off;"]