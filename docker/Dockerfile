# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as BUILDER

ARG BINARY=vision

WORKDIR /vision

# Copy of all the files to /app folder
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-gs-ping


CMD ["/vision" ]

# export PATH=$PATH:/usr/local/go/bin


# CMD ["tail", "-f", "/dev/null"]
#sudo docker run -it  096d9dbf49b4 sh