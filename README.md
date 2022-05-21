
# echogorm1

```
# MVC
presentation -call-> business -call-> model

# Clean
presentation -call-> business
model -injected-to-> business
```

## How to stasrt

```shell

docker-compose up

mysql -u root -p -h 127.0.0.01 -P 3306
CREATE DATABASE test1;

```

## Manual test example

```shell

curl -X POST -d 'email=test&password=pass' http://localhost:1323/guest/login

curl -X POST -H 'content-type: application/json' -d '{"email":"test","password":"pass"}' http://localhost:1323/guest/login


curl -X POST -d 'email=test&password=pass' http://localhost:1323/guest/register
```
