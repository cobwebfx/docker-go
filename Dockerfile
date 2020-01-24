FROM golang:1.13.4-alpine as builder
# All these steps will be cached
RUN mkdir /go/src/app
WORKDIR /go/src/app
COPY ./src/go.mod .
#COPY ./src/go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY ./src .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app



FROM nginx
EXPOSE 80
COPY --from=builder /go/src /usr/share/nginx/html
