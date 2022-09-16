# Login

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
json : {
  "email":"",
  "password":""
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
  "email":"@gmail.com",
  "password":"qwerty"
  
  
});

let reqOptions = {
  url: "http://localhost:9000/apilogin/ln",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);

```