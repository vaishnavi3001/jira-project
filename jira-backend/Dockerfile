FROM golang:1.16-alpine as builder

WORKDIR /app

COPY . .

RUN apk update \
    && apk --no-cache --update add build-base \
    && go mod download \
    && go build -o jira-backend

FROM alpine
WORKDIR /app
COPY --from=builder '/app/jira-backend' .
COPY --from=builder '/app/config' '/app/config/'
EXPOSE 6000
CMD ["./jira-backend"]



