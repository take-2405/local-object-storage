# local-object-storage
use minio

### Get Start

#### start minio
```cassandraql
docker-compose up -d
```

**Confirm**  

connect monio
```
http://localhost:9090/
```


#### start proxy
```cassandraql
go run cmd/main.go
```
default proxy is running port 8080

**EndPoints**

|  EndPoint |  Discribe  |
| ---- | ---- |
|  read/bucket  |  Check the existing bucket.  |
|  create/bucket |  Create a bucket.  |
| upload/images | Upload the image. (Only png and jpg are supported.)|