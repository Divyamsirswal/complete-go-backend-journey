## Why we use drivers (lib/pq)

* to get connected to database obs i guess
* they having inbuilt functions and features to use 


## What is db.Exec vs db.Query

* ``db.Exec`` -> i guess it use for doing creating ,updating task in sql
* ``db.Query`` -> for select kinda operations

## Why we use $1, $2

* these are senitizers i guess they are mainly for security purpose so no sql injeection can happen

## Why we use defer rows.Close()

* to ensure that database resources are released back to the connection pool regardless of how the surrounding function exits