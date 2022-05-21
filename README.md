
# echogorm1

This is example of [how to structure your golang project](https://kokizzu.blogspot.com/2022/05/how-to-structure-layer-your-golang-project.html) article with echo and gorm (you can change it to whatever framework and persistence libraries you like, the structure should still be similar).

```
# MVC
presentation -call-> business -call-> model

# Clean
presentation -call-> business
model -injected-to-> business
```

## How to start

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

## TODO
- censor logs
- add metrics
- golangci lint
- godotenv to load config from env
