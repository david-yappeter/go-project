FROM golang:1.14-alpine AS build
RUN apk add git

RUN mkdir /app
WORKDIR /app

# COPY go.mod .
# COPY go.sum .

# RUN go mod 

COPY . .

RUN go build -o /bin/myapp

CMD ["/bin/myapp"]
