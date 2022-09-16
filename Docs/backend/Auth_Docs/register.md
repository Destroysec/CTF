# Register

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
json : {
  "username":"",
  "email":"",
  "password":"",
  "repassword":""
}
```
## Example
```js
import axios from "axios";

let headersList = {
 "Accept": "*/*",
 "User-Agent": //User ,
 "X-API-KEY": //.env,
 "Content-Type": "application/json" 
}

let bodyContent = JSON.stringify({
  "username":"username",
  "email":"@gmail.com",
  "password":"qwerty",
  "repassword":"qwerty"
});

let reqOptions = {
  url: "http://localhost:9000/apilogin/reg",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);

```
[< back >](https://github.com/Destroysec/CTF/blob/main/Docs/backend/ListOfContents.md)
