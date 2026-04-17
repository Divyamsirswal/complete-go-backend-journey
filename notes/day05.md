## Why functions are better than writing code in main()?
* it can be messy writing repeated code in main so i guess you should define function for that
* its easy to debug if you make a function
* it also make your code much cleaner and better

## What is parameter vs argument?
* so what we write in main to define a function body -> the value we pass -> argument
* the function we have defined outside the main -> having some values -> parameter

## What is return type?
* so lets take a function here 
```go
    func add(a int, b int) `int` {
    	return a + b
    }
    // int -> return type , and it can be any type according to the function you have defined
```

## Why Go allows multiple return values?
* to facilitate clean, explicit error handling and to improve code readability by eliminating the need for wrapper structs or in-band error signaling