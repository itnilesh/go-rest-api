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


API Docs (Will add swagger later)

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

4. Delete Single Employee 
~~~
curl -X DELETE  http://localhost:8888/employees/{id}
~~~

5. Update Single  Employess

~~~
curl -X POST \
  http://localhost:8888/employees \
  -H 'Content-Type: application/json' \
  -d '{
		"house_number" : "#09-5A",
		"street" : "350 Ellis Street",
		"zip_code": "94043",
		"city" : "Mountain View",
		"district":"California",
		 "State" : "California",
		 "Country":"USA"
	}'
~~~
