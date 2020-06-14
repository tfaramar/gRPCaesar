const PROTO_PATH = __dirname + '/../cipher.proto';
const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
//cipher_proto variable points to cipher package, so it now has all the proto definitions
const cipher_proto = grpc.loadPackageDefinition(packageDefinition).cipher;

//Implement the service interface generated from the service definition
//Encode
const encode = (request) => {
    console.log(`Encode called with request: ${JSON.stringify(request)}`)
    let text = request.text.toLowerCase();
    let shift = request.shift;
    
    const ciphered = [];
    const newKey = shift % 26;
    const alpha = 'abcdefghijklmnopqrstuvwxyz'.split('');

    for (let letter of text) {
        if (letter === ' ') {
            ciphered.push(letter);
        } else {
            ciphered.push(getNewLetter(letter, newKey, alpha));
        }  
    }
    //return must match proto definition Message = {msg: string} 
    return {msg: ciphered.join('')};
}

const getNewLetter = (letter, key, alpha) => {
    const newLetterIdx = alpha.indexOf(letter) + key;
    return newLetterIdx <= 25 ? alpha[newLetterIdx] : alpha[-1 + (newLetterIdx % 25)];
}

//this func is passed a call obj for the RPC, which has the Input param as a property, and a cb to which we pass our returned Message
// function encodeMessage(call, callback) {
//     console.log("Encode call outputting:", encode(call.request));
//     callback(null, encode(call.request)); //first null indicates no error
// }


//Decode
const decode = (request) => {
    request.shift = (0 - request.shift);
    return encode(request);
}

// function decodeMessage(call, callback) {
//     callback(null, decode(call.request));
// }

//Handle Request
const handleRequest = (call, callback) => {
    if (call.request.encode) {
        callback(null, encode(call.request)); 
    } else {
        callback(null, decode(call.request));
    }
}


//Run the gRPC server
const getServer = () => {
    const server = new grpc.Server();
    server.addService(cipher_proto.CipherService.service, { //uses svc descriptor, which is a property of the stub
        encodeMessage: handleRequest, //implement svc methods
        decodeMessage: handleRequest
    });
    return server;
}

let routeServer = getServer(); //create server instance
routeServer.bind('0.0.0.0:50051', //specify address/port
    grpc.ServerCredentials.createInsecure());
console.log('Server listening on Port 50051');
routeServer.start();

