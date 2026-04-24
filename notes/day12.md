## Why JSON is used in backend
* for communication 
* for exchanging the data between frontend and backend

## Difference between Marshal and Unmarshal
* ``Marshal`` -> go struct -> JSON
* ``Unmarshal`` -> JSON -> go struct


## Why struct tags are important
* without this the JSON will be uneven what if u have did Name string in struct and you are accepting name tag in frontend not gonna work there should be generalize stuff so struct tags are important 

## What is r.Body
* it’s usually something like a pointer to an internal struct
