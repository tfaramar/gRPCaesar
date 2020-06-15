# gRPCaesar Microservice

## Description

gRPCaesar is a Go gRPC client that connects to a Node.js gRPC service to encode and decode text using the Caesar cipher encryption technique.

## Setting flag values

| Flag       | Description   | Value    | Default Value |
| ------------- |-------------|-------------|-------------|
| ```-t``` | text, a quote-enclosed string | ```"Et tu brute"```| ```"hello world"```
| ```-s``` | shift, an integer | ```3``` | ```0``` |
| ```-m``` | mode, a string character "e" or "d" (encode or decode) | ```e``` | ```e``` |


For example:
```bash
./grpcaesar -t "Et tu brute" -s 3 -m e
```

## Quickstart

**Start server**

```bash
cd service
npm install
node server.js
```

**Start client**

```bash
cd client
./grpcaesar -t "hi" -s 2 -m e    #encode "hi" with a shift of 2
./grpcaesar -t "jk" -s 2 -m d    #decode "jk" with a shift of 2
```

**Rebuild the Go client executable**

```bash
cd client
go build -o grpcaesar
```

**Run the Go code without executable**

```bash
cd client
go run main.go cipher.pb.go -t "hi" -s 2 -m e   #encode "hi" with a shift of 2
```

**Run Go unit tests**
```bash
cd client
go test
```
