FROM golang
WORKDIR /go/src/app
COPY . .
CMD ["/go/src/app/cicd-demo"]
