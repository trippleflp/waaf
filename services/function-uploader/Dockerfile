##
## Build
##
FROM golang:1.19.1-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /uploader

##
## Deploy
##
FROM alpine

EXPOSE 10004

WORKDIR /
COPY --from=build /uploader /uploader

CMD [ "/uploader" ]
