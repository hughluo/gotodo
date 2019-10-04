FROM golang:1.12-alpine AS build
#Install git
RUN apk add --no-cache git
#Get the hello world package from a GitHub repository
RUN go get github.com/hughluo/gotodo
WORKDIR /go/src/github.com/hughluo/gotodo
# Build the project and send the output to /bin/gotodo
RUN go build -o /bin/gotodo

FROM golang:1.12-alpine
#Copy the build's output binary from the previous build container
COPY --from=build /bin/gotodo /bin/gotodo
ENTRYPOINT ["/bin/gotodo"]