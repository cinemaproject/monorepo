FROM nginx:1.18

COPY ./static/** /usr/share/nginx/html/
COPY ./config/**  /etc/nginx/conf.d/
