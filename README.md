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
EMP_SERVICE_LOGLEVEL=panic|fatal|error|warn|info|debug
~~~

~~~
docker run  -e EMP_SERVICE_LOGLEVEL=debug -it -d -p 8888:8888  salpe/employee-service
~~~

Run command 
~~~
docker logs <docker-container-id>
~~~
Sample logs 
~~~
INFO[0000] Loading environment vars
INFO[0000] Employee service congfigs                     config="Host=localhost, Port=8888, LogLevel=debug"
INFO[0000] Starting employee service....
INFO[0000] Registerted rest end-point/employees
INFO[0000] Registerted rest end-point/employees/{id}
~~~


### Single Container Performance Test 
We need to size container. It should able to serve largest request else in auto-scaler it will just keep on creating new instances without meeting the purpose.Also we need to know how many requests per second container can sustain.

### Single User Performance Test
 This also could be done using [https://github.com/tsenart/vegeta] with single attacker config.
 In this we can increase input size progressively, till container goes out of memory or request takes very long time.

### Parallel Users Performance Test
Tool could be used is vegeta [https://github.com/tsenart/vegeta]
We need to see ideal config for CPU/Memory to meet our concurrent users SLAs.



### Metrics 
we need number fo requests per sec (throughput) and latency (time taken for single request to serve ) based on which this app will autoscale.

We need to instrument code to collect metrics.


Metrics scrapping  end-point

~~~
/internal/metrics
~~~

### Auto-scaler 
Based on metrics collected in prometheus , we will autoscale for throughput/latency numbers to meet SLAs.


### Distributed request tracing
We can add distributed request tracing using zipkin [https://zipkin.io/].


### Distributed logging 
Individual container logs need to be aggrgated and analyzed. We need tool like elastic search ELK stack [https://logz.io/learn/complete-guide-elk-stack/] for that.


### Circuit Breaker 
If say MongoDb service is down then better to trip circuit breaker then hitting the service repeatedly.
We will use Hystrix [https://github.com/Netflix/Hystrix] based circuit breaker 

### Load Balancers 
Mostly all cloud infra providers have this feature. example AWS loadbalancer. As this is HTTP service we can use HTTP based load balancer than TCP based load balancers. If we want to enable GRPC endpoints we will need TCP based loadbalancer.

### Distributed Caching 
If there is demand for particular results many time or result is compute heavy we can store in distributed cache like Redis [https://redis.io/]. Before serving request, we need to check in cache if it is there we need to return from cache else re-compoute results.
Cache is good for mostly read only data. But if there are changes to data we need to invalidate cache.

### Asyc processing for large input 
Asyc processing in REST api is quiet tricky .

If request can not be served in real time due to large processing time or need special resources to serve such requests we need async processing framework or job framework. Example we need to export details of all employees in CSV file.
Such arragement needs REST apis designed for async processing using location header and redirects.

Design of Asyc rest api 

Approach ONE :

1. Submit async request 
2. Return with 202 Accpeted header and  job url in location header where user can check status of job and give expected time to poll.
3. User will call job url in intervals specified and he will get status 
4. Once job gets successfull , call to job status url (302 Redirect) will redirect to actual results.

Approach TWO:

1. Submit async job request with call-back endpoint. Call-back url has pre-defined API singnature.
2. When async job completed, it will call call-back url with details of completed jobs.


Flow for distributed processing once server gets request 
This is producer-consumer pattern , can be impleted using queuing framework like RabbitMQ [https://www.rabbitmq.com/]

1. User submits asyc job request.
2. Convert request to queue item and persist in DB.
3. Put item representing job is put in queue.
4. Worker corresponding to job type picks up and processed.
5. There should be some retry logic like failure queue /retry queue / success queue. You can use dead-letter routing technic available in rabbitMQ
6. Persist results in DB
7. Deque item from queue
8. Notify user about result 


### Authentication And Authorization 
We some kind of authorization, nowadays OAuth2.0 is used widely. May be passing JWT token passed thought `Authorization` header and `Bearer` scheme could be OK.
We need to build `authentication and authorization server` so that user can get token based on their username/password then same token will be used for all API communincation.
This token will expire after some time say in 24 hours then user need to renew/re-issue the toke.

For authorization , we need to define roles in token and in code we need to check those before calling psrticualr action.
This can be easily done using spring security and bunch of annotations.






















