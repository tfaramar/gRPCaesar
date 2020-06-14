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
    console.log('About to return', ciphered.join(''));
    return ciphered.join('');
}

const getNewLetter = (letter, key, alpha) => {
    const newLetterIdx = alpha.indexOf(letter) + key;
    return newLetterIdx <= 25 ? alpha[newLetterIdx] : alpha[-1 + (newLetterIdx % 25)];
}

function encodeMessage(call, callback) {
    // console.log("FROM ENCODE MSG CALL", encode(call.request));
    callback(null, encode(call.request));
}


//Decode
const decode = (request) => {
    request.shift = (0 - request.shift);
    return encode(request);
}

function decodeMessage(call, callback) {
    callback(null, decode(call.request));
}


//Build server
const getServer = () => {
    const server = new grpc.Server();
    server.addService(cipher_proto.CipherService.service, {
        encodeMessage: encodeMessage,
        decodeMessage: decodeMessage
    });
    return server;
}

let routeServer = getServer();
routeServer.bind('0.0.0.0:50051',
    grpc.ServerCredentials.createInsecure());
console.log('Server listening on Port 50051');
routeServer.start();

