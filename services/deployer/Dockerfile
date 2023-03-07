##
## Build
##
FROM golang:1.19.1-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build -o /deployer

##
## Deploy
##
FROM alpine

EXPOSE 10003

WORKDIR /
COPY --from=build /deployer /deployer

CMD [ "/deployer" ]
