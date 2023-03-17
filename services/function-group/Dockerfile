##
## Build
##
FROM golang:1.19.1-alpine AS build
WORKDIR /app/service/api

COPY /libs /app/libs
COPY /api-gateway /app/service/api-gateway
COPY /deployer /app/service/deployer

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /fn

##
## Deploy
##
FROM alpine

EXPOSE 10000

WORKDIR /
COPY --from=build /fn /fn

CMD [ "/fn" ]
