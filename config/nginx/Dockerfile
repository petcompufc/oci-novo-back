FROM nginx:1.25.3-alpine

COPY nginx.conf mime.types /etc/nginx/

RUN mkdir /etc/nginx/logs
RUN touch /etc/nginx/logs/error.log

EXPOSE 80
