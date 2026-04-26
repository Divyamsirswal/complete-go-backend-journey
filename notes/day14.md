## Why use switch for methods instead of if ?
* ``switch`` -> i guess switch way looks much cleaner than the if else ladder

## Why JSON encoding directly to w ?
* because ``w`` is the response writer which we send to the client from where the request came so this ``w`` is where JSON encode directly

## Why validation is important ?
* for the security purpose
* integrity in the data

## What problems with in-memory storage ?
* this is only for the temporary purpose 
* only uses when you are developing api's and testing out things
* because after closing the server all data removes (as data is in the RAM)