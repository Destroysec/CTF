# Set_Github

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
    jwt: #JWT's Login or Register
json : {
  "github":"https://github.com/Ax-47"
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
}

let formdata = new FormData();
formdata.append("github", "https://github.com/Ax-47");

let bodyContent =  formdata;

let reqOptions = {
  url: "localhost:9000/setGithub",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);
```
[< back >](https://github.com/Destroysec/CTF/blob/main/Docs/backend/ListOfContents.md)
