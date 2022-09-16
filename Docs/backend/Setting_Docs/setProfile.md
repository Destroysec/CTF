# Set_Profile

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
    jwt: #JWT's Login or Register
json : {
  "file": C:\Users\img\profile.PNG,
 
}
```
## Example
```js
import axios from "axios";
var fs = require('fs');

let headersList = {
 "Accept": "*/*",
 "User-Agent": //User ,
 "X-API-KEY": //.env,
 "jwt": //JWT's Login or Register,
}

let formdata = new FormData();
formdata.append("file", fs.createReadStream("C:\Users\img\profile.PNG"));

let bodyContent =  formdata;

let reqOptions = {
  url: "http://localhost:9000/setProfile",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);
```
[< back >](https://github.com/Destroysec/CTF/blob/main/Docs/backend/ListOfContents.md)