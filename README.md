# gRPCaesar Microservice

gRPCaesar is a Go gRPC client that connects to a Node.js gRPC service to encode and decode text using the Caesar cipher encryption technique.




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
