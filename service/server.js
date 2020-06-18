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

const cipher_proto = grpc.loadPackageDefinition(packageDefinition).cipher;


//Encode
const encode = (request) => {
    //console.log(`Encode called with request: ${JSON.stringify(request)}`)
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

    return {msg: ciphered.join('')};
}

const getNewLetter = (letter, key, alpha) => {
    const newLetterIdx = alpha.indexOf(letter) + key;
    return newLetterIdx <= 25 ? alpha[newLetterIdx] : alpha[-1 + (newLetterIdx % 25)];
}

//Decode
const decode = (request) => {
    request.shift = (0 - request.shift);
    return encode(request);
}

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
    server.addService(cipher_proto.CipherService.service, { 
        processMessage: handleRequest
    });
    return server;
}

let routeServer = getServer(); 
routeServer.bind('0.0.0.0:50051', 
    grpc.ServerCredentials.createInsecure());
console.log('Server listening on Port 50051');
routeServer.start();

