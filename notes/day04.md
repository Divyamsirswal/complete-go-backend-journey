## Why use nested loops instead of single loop?
* I think when we have matrix kinda structure and we want to go through it and do some operation we have to use nested loops
* when you have multiple dimensions you have to use nested loop , single loop not gonna help

## Difference between normal loop and range
* ```normal loop``` -> Its a traditional way to access the data works well for most of the things but you should have the clear understanding of the types and all

* ```range``` -> generally these range loops used by certain predefined data structures, it give ease to you that you can have the value as well as index. These ```range loop``` mostly used in the backend dev, like when you want to take data from the db you dont the length but you have the structure.

## What is string indexing?
* ```indexing``` means you trying to index that particular block from the memory you have defind 
    >  b =  [v1 v2 v3 v4 v5]
    >   v1 - 0, v2 - 1, v3 - 2, v4 - 3, v5 - 4
* when you want to access the data you have in the data structure you have to use indexing for sure.
     *  value = b[0] -> v1 like that 
* talking about string indexing its same (golang has something different for the string)
    * str = "test" -> its string of ascii values when u try to access them
        * str[0] -> ascii value of ```t```
        * you have to use either ```string(str)``` or convert the each str char using ```rune```
           
## Why do we convert string(name[i])?

* the thing is that ```name := "test"``` it is string but when u try to access ```name[i]``` it is in ascii so you have to convert back then complete ```name -> string(name)```  
