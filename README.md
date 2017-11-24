# Still needs

* [ ] presenters
* [ ] better generators
* [ ] better generators
* [ ] migration runners
* [ ] migration generators date +"%Y%m%d%H%M%S"



## Creating a grpc client

```
export LD_LIBRARY_PATH=/usr/local/lib
cd /opt && git clone https://github.com/grpc/grpc.git
cd grpc && make
creating a php client
protoc --proto_path=app/blog/proto --php_put=php_things --grpc_out=php_things --plugin=protoc-gen-grpc=/opt/grpc/bins/opt/grpc_php_plugin ./app/blog/proto/blog.proto
```
