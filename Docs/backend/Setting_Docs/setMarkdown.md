# Set_MarkDown

## how to requests

## info 
```yml
method : post
header :
    X-API-KEY : #.env 
    jwt: #JWT's Login or Register
json : {
  "input": "# test", #char<256
 
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
formdata.append("input", "# test");

let bodyContent =  formdata;

let reqOptions = {
  url: "localhost:9000/setMarkdown",
  method: "POST",
  headers: headersList,
  data: bodyContent,
}

let response = await axios.request(reqOptions);
console.log(response.data);
```
[< back >](https://github.com/Destroysec/CTF/blob/main/Docs/backend/ListOfContents.md)