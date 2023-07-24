# FROM golang:1.20-alpine as builder
# RUN mkdir /app
# COPY . /app
# WORKDIR /app
# RUN go build -o brokerApp ./cmd/api
# RUN chmod +x /app/brokerApp


FROM alpine:latest
RUN mkdir /app
COPY loggerServiceApp /app
CMD [ "/app/loggerServiceApp" ]