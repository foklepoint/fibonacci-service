# fibonacci-service

### Simple service to calculate fibonacci

#### Running 
```bash
$ go get github.com/foklepoint/fibonacci-service/
$ fibonacci-service& # Run our service in the background
$ curl -X POST localhost:8080/calculate -d '{ "nth": 12 }'
# Kill the service using kill -9 %
```


###
 
 #### Building a minimal docker binary
 ```bash
 $ make
 $ docker run -p 8080:8080 fibonacci-service
 $ curl -X POST localhost:8080/calculate -d '{ "nth": 12345 }'
 ```
