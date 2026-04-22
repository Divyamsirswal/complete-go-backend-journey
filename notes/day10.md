## Why Go uses error instead of exceptions ?
* error used as an value here in go so, go always ensure you are returing value + error always so it makes the developer flow very much smoother to go on.
* there is no try catch concept here in go so you have to use error manually i guess

## What does err != nil mean ?
* let say you are returing something from the function + error also and if the thing are good you will return the value + nil which means no error but if there is an error you will return some_value + error -> now error is not nil i hope you get it , its simple and you can check this always after the function call -> go way of writing code

## Why ignoring errors is dangerous ?
* if you don't handle errors that can create lots of confusion while developing smthing like solo or team.
* if you not properly know what is error you dont know the other cases of that particular part
* thats why go always uses errors as value and you have to handle by yourself and its really better i can say that.

## Difference between error and panic (basic idea) ?
* ``error`` -> some unusual occur during programming.
* ``panic`` -> something crashes out of the box