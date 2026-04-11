## What is a variable ?
- In simple term it's a container which hold the data.
- think like a small box where u can put items and use them  later according to the use case you have.


## Difference between *var* and *:=*
- ```var``` can be used with OR without value
   - ```var Name type ```
-  generally we use this also when u want to export an variable to package level, which become really important when you do backend development.

- ```:=```  you gonna use this 90% of the time in go, but this is only scope to function level you can't declare it outside the function keep in mind. can be used only with value keep in mind
    - ```name := "david" ```


## What is a data type?
- A data type defines what kind of value a variable can store and what operations can be performed on it
    - var Name ***type*** (explicit)
    - name := "test" (implicit)
    - Types
        - int, bool, float
        - struct, interface  


## Why Go is strict about types?
- Go is strict so your code is safe, fast, and predictable.
- prevents bugs at compile time
- avoids unexpected behavior
- improves performance (no runtime guessing)