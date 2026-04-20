## Why use nested structs?
* when we want to store data inside data like metadata so lets take an example if you have user info + you also want the address part one option for single struct but what if address has some other data also. i guess nested struct can make life easier with that.

## Difference between struct and nested struct
* ```struct```
    * ```go
        type User struct{
            Name string
            Age int
            Address string 
        }
        ```
    * ```go
        type Address struct{
            City string
            State string
        }

        type User struct{
            Name string
            Age int
            Address Address
        }
      ```

## Why use pointer in struct functions?
* we use pointers in struct functions because when you are working with data you have an option to pass by value but that can create big issue with your ram cost and make things harder at scale , thats why use pointers -> simply no copy -> use original 

## What are struct tags used for?

* go uses it for serialize and deserialize your data while doing client server communication.
* ```go
    type User struct{
        Name string `json:"name"`
        Age int     `json:"age"`
    }
  ```