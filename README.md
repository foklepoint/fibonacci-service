# fibonacci-service
---
#### Simple service to calculate fibonacci

```bash
$ go get foklepoint/fibonacci-service/
$ fibonacci-service& # Run our service in the background
$ curl -X POST localhost:8080/calculate -d '{ "nth": 12 }'
# Kill the service using kill -9 %
```
