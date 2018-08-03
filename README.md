# go-rest-API

### Objective

Create a simple tutorial for beginners,  who want to write rest based micro-services in Go. 

### How to run?

Pull image from https://hub.docker.com/

~~~
docker pull salpe/employee-service

docker run  -it -d -p 8888:8888  salpe/employee-service

~~~

To build image locally on Mac

1. Set up Go [https://golang.org/doc/install],  GOROOT and GOPATH
2. Run the following command
~~~
go get github.com/itnilesh/go-rest-api

~~~

3.  Run the following command
  ~~~
    cd $GOPATH/src/github.com/itnilesh/go-rest-api 
    
    dep ensure 
   ~~~
4.  Run the following command
  ~~~
  docker build . 
  ~~~
5. Run following command
~~~~
docker tag <image-id>  salpe/employee-service
~~~~

6. Run the following command
~~~ 
docker run  -it -d -p 8888:8888  salpe/employee-service 
~~~
7. To test application is up or not, you should get an empty collection
~~~
curl -X GET  http://localhost:8888/employees 
~~~


API Docs 

Please refer swagger docs here 
https://github.com/itnilesh/go-rest-api/blob/master/docs/api/employee-service.yaml

Simple curl commands 

1. Create employee 
~~~
curl -X POST \
  http://localhost:8888/employees \
  -H 'Content-Type: application/json' \
  -d '{
    "first_name": "nilesh",
    "last_name":"salpe",
    "address": {
        "house_number" : "#09-99",
        "street" : "Upper Serangoon Road",
        "zip_code": "536564",
        "city" : "Singapore",
        "district":"Singapore",
         "State" : "Singapore",
         "Country":"Singapore"
    }
}'
~~~
2. List employess 

~~~
curl -X GET  http://localhost:8888/employees 
~~~

3.  List single employee

~~~
curl -X GET  http://localhost:8888/employees/{id}
~~~


4. Update Single  Employees

~~~
curl -X PUT \
  http://localhost:8888/employees\{id} \
  -H 'Content-Type: application/json' \
  -d '{
        "house_number" : "#09-5A",
        "street" : "350 Salpe Street",
        "zip_code": "94043",
        "city" : "Mountain View",
        "district":"California",
         "State" : "California",
         "Country":"USA"
    }'
~~~

5. Delete Single Employee 
~~~
curl -X DELETE  http://localhost:8888/employees/{id}
~~~

### Debugging 

You can set envirnment variable while starting container 
~~~
EMP_SERVICE_LOG_LEVEL=panic|fatal|error|warn|info|debug
~~~

~~~
docker run  -e EMP_SERVICE_LOG_LEVEL=debug -it -d -p 8888:8888  salpe/employee-service
~~~

Run command 
~~~
docker logs <docker-container-id>
~~~
Sample logs 
~~~
INFO[0000] Loading environment vars
INFO[0000] Employee service congfigs                     config="Host=localhost, Port=8888"
INFO[0000] Starting employee service....
INFO[0000] Registerted rest end-point/employees
INFO[0000] Registerted rest end-point/employees/{id}
~~~
