# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev
LABEL MAINTAINER "https://github.com/itnilesh"

RUN echo $GOROOT
RUN echo $GOPATH
WORKDIR /app

ENV SRC_DIR=/${GOPATH}/src/github.com/itnilesh/go-rest-api
ADD . $SRC_DIR
RUN cd $SRC_DIR; go build  ./cmd/employee-service/; cp employee-service /app/

EXPOSE 8888
ENTRYPOINT ["./employee-service"]