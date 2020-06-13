const grpc = require('grpc');
const protoLoader = require('@grpc/proto-loader');

// const cipherProto = grpc.load('../cipher.proto');
// grpc.load deprecated, so the below options are suggested
const packageDefinition = protoLoader.loadSync(
    ('../cipher.proto'),
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    });
const cipherservice = grpc.loadPackageDefinition(packageDefinition).cipherservice;

//Encode
const encodeText = (text, shift) => {
    text = text.toLowerCase();
    let ciphered = [];
    let newKey = shift % 26;
    for (let letter of text) {
        ciphered.push(getNewLetter(letter, newKey));
    }
    return ciphered.join('');
}

const getNewLetter = (letter, key) => {
    let newLetterUnicode = letter.charCodeAt() + key;
    return newLetterUnicode <= 122 ? String.fromCharCode(newLetterUnicode) : String.fromCharCode(96 + (newLetterUnicode % 122));
}


//Decode
const decodeText = (text, shift) => {
    let decodeShift = (0 - shift);
    return encodeText(text, decodeShift);
}

let test = encodeText("hi this is a friend", 14)
console.log(test);

console.log(decodeText(test, 14));



const server = new grpc.Server();

//REFACTOR THIS:
// server.addService(cipherProto.cipher.CipherService.service, {});

// server.bind('0.0.0.0:50051', 
//     grpc.ServerCredentials.createInsecure());
// console.log('Server running on Port 50051');
// server.start();