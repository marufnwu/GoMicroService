FROM alpine:latest
RUN mkdir /app
COPY frontApp /app
EXPOSE 8082
CMD [ "/app/frontApp" ]