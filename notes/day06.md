## Difference between array and slice
* ```array``` -> are of fixed size used very less when you are working in real systems
* ```slice``` -> are of dynamic size and have better performance then normal array, used most in the real systems

## Why slices are used more in Go
* i guess because they are dynamic and more realiable to use and efficiency is far better than normal fixed array

## What does append() do internally (basic idea)
* i guess you are asking it for the ```slice```
    * Case -> when you have enough space
        * it just add the element at the end simply
    * Case -> when you dont have enough space
        * it create a new array of double size generally and copy the rest of the elements to it and do this process if this case hits

## Difference between len and cap
* ```len``` -> the number of elements are there in the slice
* ```cap``` -> the actual reserved size of the slice 