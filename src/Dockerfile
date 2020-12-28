FROM golang AS build
ARG sha
ARG ver
ARG date
WORKDIR /go/src
COPY registry.go .
RUN go get github.com/gin-gonic/gin
RUN go get github.com/golang/glog
RUN go build -o registry -ldflags "-X main.sha=$sha -X main.version=$ver -X 'main.date=$date'"

FROM centos
EXPOSE 5000
COPY --from=build /go/src/registry .
ENTRYPOINT ["./registry"]
CMD ["-stderrthreshold=INFO"]
