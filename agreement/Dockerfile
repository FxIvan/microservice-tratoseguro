# base image
FROM golang:1.19.3-alpine AS builder
# create appuser.
ARG MONGO_USER_VAR
ARG MONGO_PASSWORD_VAR
RUN adduser -D -g '' elf
# create workspace
WORKDIR /opt/app/
COPY go.mod go.sum ./
# fetch dependancies
RUN go mod download && \
    go mod verify
# copy the source code as the last step
COPY . .
# build binary
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/microservicio-agreement-api ./cmd/app


# build a small image
FROM alpine:3.17.3
LABEL language="golang"
LABEL org.opencontainers.image.source https://github.com/FxIvan/microservice-tratoseguro
# import the user and group files from the builder
COPY --from=builder /etc/passwd /etc/passwd
# copy the static executable
COPY --from=builder --chown=elf:1000 /go/bin/microservicio-agreement-api /microservicio-agreement-api
# use a non-root user
USER elf

ENV MONGO_USER=MONGO_USER_VAR
ENV MONGO_PASSWORD=MONGO_PASSWORD_VAR

# run app
ENTRYPOINT ["./microservicio-agreement-api"]

#docker build --build-arg MONGO_USER_VAR=admtratoseguro210814 --build-arg MONGO_PASSWORD_VAR=LkdU7ZDADARiFEtZiKJUjUeg5Swfyq9dA7qwkqjerkpQZwEvUs  -t almendraivan/microservice-agreement:v0 .
#docker run  -p 9092:4000 --name codigo-go-agreement almendraivan/microservice-agreement:v0