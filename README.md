# install protobuf compiler


## download & install

```
wget https://github.com/google/protobuf/releases/download/v3.11.2/protobuf-all-3.11.2.zip
unzip protobuf-all-3.11.2.zip && cd protobuf-3.11.2/
./configure
make
make install
```
## check version

```
ldconfig
protoc --version
```

## go grpc install

```
go get -u google.golang.org/grpc
```

## Protoc Plugin

```
go get -u github.com/golang/protobuf/protoc-gen-go
```

# compile proto

```
protoc --go_out=plugins=grpc:. ./proto/*.proto
```


# install grpcurl
```
go get github.com/fullstorydev/grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl
```
