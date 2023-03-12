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
FROM buildah/buildah

EXPOSE 10004

ARG USER=podman
ENV HOME /home/$USER

# install sudo as root
RUN apk add --update sudo
# add new user
RUN echo "$USER ALL=(ALL) NOPASSWD: ALL" > /etc/sudoers.d/$USER && chmod 0440 /etc/sudoers.d/$USER



WORKDIR $HOME

COPY --from=build /uploader /uploader

CMD [ "/uploader" ]
