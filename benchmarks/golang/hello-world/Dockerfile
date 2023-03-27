##
## Build
##
FROM golang:1.19.1-alpine AS build
WORKDIR /app

#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download

COPY . ./
RUN go build -o helloworld

##
## Deploy
##
FROM alpine

WORKDIR /
COPY --from=build /app/helloworld .

CMD [ "/helloworld" ]
