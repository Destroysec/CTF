# OTP

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
    jwt: #JWT's Login or Register
json : {
  "jwt": #JWT's Login or Register
  "otp": "otp"
}
```
## Example
```js
import axios from "axios";

let headersList = {
 "Accept": "*/*",
 "User-Agent": //User ,
 "X-API-KEY": //.env,
 "jwt": //JWT's Login or Register,
 "Content-Type": "application/json" 
}

let bodyContent = JSON.stringify({
  "jwt": //JWT's Login or Register,
  "otp": "otp"
});

let reqOptions = {
  url: "http://localhost:9000/verifyotp",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);
```
[< back >](https://github.com/Destroysec/CTF/blob/main/Docs/backend/ListOfContents.md)