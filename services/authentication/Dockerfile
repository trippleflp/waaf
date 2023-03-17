##
## Build
##
FROM golang:1.19.1-alpine AS build
WORKDIR /app/service/api

COPY /libs /app/libs

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /api

##
## Deploy
##
FROM alpine

EXPOSE 10000

WORKDIR /
COPY --from=build /api /api

CMD [ "/api" ]
