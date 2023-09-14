FROM golang:1.19-alpine AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o main .


#final stage
FROM scratch
COPY --from=builder ["/build/main", "/build/.env", "/"]
ENTRYPOINT ["/main"]

CMD ["./main"]






#FROM golang:1.16
#
## Set the working directory inside the container
#WORKDIR /app
#
## Copy the local package files and Go modules to the container's workspace
#COPY go.mod go.sum ./
#RUN go mod download
#
## Copy the rest of your application code to the container
#COPY . .
#
## Build the Go application inside the container
#RUN go build -o main
#
## Expose the port your application will run on (change as needed)
#EXPOSE 8000
#
## Define the command to run your application
#CMD ["./main"]
