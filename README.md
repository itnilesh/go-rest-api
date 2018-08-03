# go-rest-api
Sample rest apis in GO

~~~
docker pull salpe/employee-service

docker run  -it -d -p 8888:8888  salpe/employee-service

~~~

To build image lcoally 

1. Set up Go GOROOT and GOPATH
2. go get `github.com/itnilesh/go-rest-api`
3. 
  ~~~
    cd $GOPATH/src/github.com/itnilesh/go-rest-api 
   ~~~
4. 
  ~~~
  docker build . 
  ~~~
5. 
~~~~
docker tag <image-id>  salpe/employee-service
~~~~

6. 
~~~ 
docker run  -it -d -p 8888:8888  salpe/employee-service 
~~~

