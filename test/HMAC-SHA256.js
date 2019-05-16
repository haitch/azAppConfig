var sha256 = require("crypto-js/sha256");
var hmacSHA256 = require("crypto-js/hmac-sha256");
var Base64 = require("crypto-js/enc-base64");

function signRequest(host, 
    method,      // GET, PUT, POST, DELETE
    url,         // path+query
    timestamp,   // timestamp
    body,        // request body (undefined of none)
    credential,  // access key id
    secret)      // access key value (base64 encoded)
{
var verb = method.toUpperCase();
var utcNow = timestamp;
var contentHash = Base64.stringify(sha256(body));

//
// SignedHeaders
var signedHeaders = "x-ms-date;host;x-ms-content-sha256"; // Semicolon separated header names

//
// String-To-Sign
var stringToSign = 
verb + '\n' +                              // VERB
url + '\n' +                               // path_and_query
utcNow + ';' + host + ';' + contentHash;   // Semicolon separated SignedHeaders values

//
// Signature
var signature = Base64.stringify(hmacSHA256(stringToSign, Base64.parse(secret)));

//
// Result request headers
return [
{ name: "x-ms-date", value: utcNow },
{ name: "x-ms-content-sha256", value: contentHash },
{ name: "Authorization", value: "HMAC-SHA256 Credential=" + credential + ", SignedHeaders=" + signedHeaders + ", Signature=" + signature }
];
}

console.log(signRequest("haitchgo.azconfig.io", "GET", "/keys", "Tue, 14 May 2019 07:36:46 GMT", "", "0-l1-s0:Et/0nCwtWNbYBHodqARK", "UQeiC2Ln7XTK11uPHEz5WWIQutoZalP99nZSm8fyWiU="))