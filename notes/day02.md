## Why does ```Scan()``` needs ```&``` ?
* ```fmt.Scan()``` needs the memory address of a variable so it can directly modify its value.

## Difference between Print, Println, Printf
* ```Print``` -> no new line only print at point 
* ```Println``` -> print and add a new line after it
* ```Printf``` -> Printf does NOT add newline automatically, you must use ```\n```

## What happens if user enters "abc" instead of number?
* Go will fail to scan properly
* It may leave variable unchanged OR set zero value
* AND return an error (which you are currently ignoring)

## Why is formatting important in real applications?
* ensures structured output
* helps in logging, APIs, debugging
* avoids confusion in large systems