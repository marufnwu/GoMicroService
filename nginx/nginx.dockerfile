# Use separate stage for deployable image
FROM staticfloat/nginx-certbot:latest

COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80