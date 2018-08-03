# go-rest-api

Sample rest apis in GO

### How to run ?

Pull image from https://hub.docker.com/

~~~
docker pull salpe/employee-service

docker run  -it -d -p 8888:8888  salpe/employee-service

~~~

To build image locally on Mac

1. Set up Go [https://golang.org/doc/install],  GOROOT and GOPATH
2. Run following command
~~~
go get github.com/itnilesh/go-rest-api

~~~

3.  Run following command
  ~~~
    cd $GOPATH/src/github.com/itnilesh/go-rest-api 
   ~~~
4.  Run following command
  ~~~
  docker build . 
  ~~~
5. Run following command
~~~~
docker tag <image-id>  salpe/employee-service
~~~~

6. Run following command
~~~ 
docker run  -it -d -p 8888:8888  salpe/employee-service 
~~~
7. Test application is up or not , you should get empty collection
~~~
curl -X GET  http://localhost:8888/employees 
~~~

